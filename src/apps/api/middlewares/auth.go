package middlewares

import (
	"net/http"
	"strings"

	"github.com/ahmetb/go-linq"
	"github.com/google/uuid"

	"eletronic_point/src/apps/api/apimsg"
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/utils"
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/utils/tokenextractor"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var (
	AuthModelPath  string
	AuthPolicyPath string

	method string
	path   string
	roles  []string
)

var logger = Logger()
var authService = dicontainer.AuthServices()

func negativeTokenCookie() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = handlers.COOKIE_TOKEN_NAME
	cookie.Path = "/"
	cookie.HttpOnly = false
	cookie.MaxAge = -1
	return cookie
}

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	enforcer, err := casbin.NewEnforcer(AuthModelPath, AuthPolicyPath)
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

var allowList = []string{
	"/api/auth/refresh",
}

func isPathInAllowList(path string) bool {
	for _, p := range allowList {
		if path == p {
			return true
		}
	}
	return false
}

func isAnAnonymousRequest() bool {
	return linq.From(roles).Contains(tokenextractor.AnonymousRole)
}

func isAccessGranted(enforcer *casbin.Enforcer) bool {
	for _, role := range roles {
		if granted, err := enforcer.Enforce(role, path, method); granted && err == nil {
			return true
		} else if err != nil {
			log.Fatal().Err(err)
		}
	}
	return false
}

func sendUnavailableServiceResponse(c echo.Context) error {
	return c.JSON(http.StatusServiceUnavailable, response.ErrorMessage{
		Code:    http.StatusServiceUnavailable,
		Message: apimsg.AuthServerUnavailable,
	})
}

func sendUnauthorizedResponse(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, response.ErrorMessage{
		Code:    http.StatusUnauthorized,
		Message: apimsg.UnauthorizedErrMsg,
	})
}

func sendForbiddenResponse(c echo.Context) error {
	return c.JSON(http.StatusForbidden, response.ErrorMessage{
		Code:    http.StatusForbidden,
		Message: apimsg.ForbiddenErrMsg,
	})
}
