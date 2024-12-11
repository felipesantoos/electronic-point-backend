package middlewares

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares/permissions"
	"eletronic_point/src/apps/api/utils"
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	jsonadapter "github.com/casbin/json-adapter/v2"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var logger = Logger()
var authService = dicontainer.AuthServices()
var permissionsHelper = permissions.New()
var casbinModelTemplate = `
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == p.sub && regexMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`

func negativeTokenCookie() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = handlers.COOKIE_TOKEN_NAME
	cookie.Path = "/"
	cookie.HttpOnly = false
	cookie.MaxAge = -1
	return cookie
}

func GuardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	enforcer, err := newCasbinEnforcer()
	if err != nil {
		log.Fatal().Err(err)
	}
	return func(ctx echo.Context) error {
		tokenCookie, err := ctx.Cookie(handlers.COOKIE_TOKEN_NAME)
		var accRole string = role.ANONYMOUS_ROLE_CODE
		var claims *authorization.AuthClaims
		if err == nil {
			token := tokenCookie.Value
			if valid, _ := sessionIsValidWith(token); !valid {
				ctx.SetCookie(negativeTokenCookie())
			} else {
				var ok bool
				if accRole, ok = utils.ExtractAuthorizationAccountRole(token); !ok {
					return ctx.NoContent(http.StatusUnauthorized)
				}
				claims, _ = utils.ExtractTokenClaims(token)
				ctx.Set("authenticated", true)
			}
		}
		method := ctx.Request().Method
		path := ctx.Request().URL.Path
		if ok, err := enforcer.Enforce(strings.ToLower(accRole), path, method); err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		} else if accRole == role.ANONYMOUS_ROLE_CODE && !ok {
			return ctx.NoContent(http.StatusUnauthorized)
		} else if !ok {
			var logData map[string]interface{} = map[string]interface{}{
				"path":   path,
				"method": method,
				"role":   accRole,
			}
			if claims != nil {
				logData["user_id"] = claims.AccountID
			}
			logger.Warn().Fields(logData).Msg("FORBIDDEN ACCESS")
			return ctx.NoContent(http.StatusForbidden)
		}
		return next(ctx)
	}
}

func newCasbinEnforcer() (*casbin.Enforcer, error) {
	authModel, err := model.NewModelFromString(casbinModelTemplate)
	if err != nil {
		return nil, err
	}
	authAdapter, err := newCasbinJSONAdapter()
	if err != nil {
		return nil, err
	}
	enforcer, err := casbin.NewEnforcer(authModel, authAdapter)
	if err != nil {
		fmt.Println("Error when building enforcer:", err)
		return nil, err
	}
	return enforcer, nil
}

func authCasbinPolicies() []map[string]string {
	authPolicies := permissionsHelper.AuthPolicies()
	policies := []map[string]string{}
	for _, policy := range authPolicies {
		policies = append(policies, map[string]string{
			"PType": "p",
			"V0":    policy.Subject(),
			"V1":    policy.Object(),
			"V2":    policy.Action(),
		})
	}
	return policies
}

func newCasbinJSONAdapter() (*jsonadapter.Adapter, error) {
	authPolicy := authCasbinPolicies()
	authPolicyBytes, err := json.Marshal(&authPolicy)
	if err != nil {
		return nil, err
	}
	authAdapter := jsonadapter.NewAdapter(&authPolicyBytes)
	return authAdapter, nil
}

func sessionIsValidWith(authToken string) (bool, errors.Error) {
	if claims, err := utils.ExtractTokenClaims(authToken); err != nil {
		return false, nil
	} else if uID, err := uuid.Parse(claims.AccountID); err != nil {
		return false, nil
	} else if exists, err := authService.SessionExists(&uID, authToken); err != nil {
		return false, err
	} else if !exists {
		return false, nil
	}
	return true, nil
}
