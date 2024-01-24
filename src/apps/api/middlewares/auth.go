package middlewares

import (
	"backend_template/src/apps/api/dicontainer"
	"backend_template/src/apps/api/handlers/dto/response"
	"backend_template/src/apps/api/middlewares/permissions"
	"backend_template/src/apps/api/utils"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/role"
	"encoding/json"
	"fmt"
	"strings"

	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	jsonadapter "github.com/casbin/json-adapter/v2"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var logger = Logger()
var authService = dicontainer.AuthUseCase()
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

func GuardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	enforcer, err := newCasbinEnforcer()
	if err != nil {
		log.Fatal().Err(err)
	}
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		method := ctx.Request().Method
		path := ctx.Request().URL.Path
		if accRole, ok := utils.ExtractAuthorizationAccountRole(authHeader); !ok {
			return response.ErrorBuilder().NewUnauthorizedError()
		} else if ok, err = enforcer.Enforce(strings.ToLower(accRole), path, method); err != nil {
			return response.ErrorBuilder().NewInternalServerError()
		} else if accRole == role.ANONYMOUS_ROLE_CODE && !ok {
			return response.ErrorBuilder().NewUnauthorizedError()
		} else if !ok {
			claims, _ := utils.ExtractTokenClaims(authHeader)
			logger.Warn().Fields(map[string]interface{}{
				"path":    path,
				"method":  method,
				"role":    accRole,
				"user_id": claims.AccountID,
			}).Msg("FORBIDDEN ACCESS")
			return response.ErrorBuilder().NewForbiddenError()
		} else if accRole != role.ANONYMOUS_ROLE_CODE {
			_, authToken := utils.ExtractToken(authHeader)
			if valid, _ := sessionIsValidWith(authToken); !valid {
				return response.ErrorBuilder().NewUnauthorizedError()
			}
			ctx.Set("authenticated", true)
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
