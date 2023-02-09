package middlewares

import (
	"dit_backend/src/api/dicontainer"
	"dit_backend/src/api/handlers/dto/response"
	"dit_backend/src/core/domain/authorization"
	"dit_backend/src/core/domain/errors"
	"dit_backend/src/core/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	casbin "github.com/casbin/casbin/v2"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var logger = Logger()
var authService = dicontainer.AuthUseCase()

func GuardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	enforcer, err := newCasbinEnforcer()
	if err != nil {
		log.Fatal().Err(err)
	}

	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		method := ctx.Request().Method
		path := ctx.Request().URL.Path
		if role, ok := utils.ExtractAuthorizationAccountRole(authHeader); !ok {
			return ctx.NoContent(http.StatusUnauthorized)
		} else if ok, err = enforcer.Enforce(strings.ToLower(role), path, method); err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		} else if role == authorization.ANONYMOUS_ROLE_CODE && !ok {
			return ctx.NoContent(http.StatusUnauthorized)
		} else if !ok {
			logger.Warn().Fields(map[string]interface{}{
				"path":   path,
				"method": method,
				"role":   role,
			}).Msg("FORBIDDEN ACCESS")
			return ctx.NoContent(http.StatusForbidden)
		} else if role != authorization.ANONYMOUS_ROLE_CODE {
			_, authToken := utils.ExtractToken(authHeader)
			if valid, err := sessionIsValidWith(authToken); !valid {
				if err != nil {
					return ctx.JSON(http.StatusUnauthorized, response.NewErrorFromCore(err, http.StatusUnauthorized))
				}
				return ctx.NoContent(http.StatusUnauthorized)
			}
		}
		return next(ctx)
	}
}

func newCasbinEnforcer() (*casbin.Enforcer, error) {
	authModel := os.Getenv("SERVER_CASBIN_AUTH_MODEL")
	authPolicy := os.Getenv("SERVER_CASBIN_AUTH_POLICY")
	enforcer, err := casbin.NewEnforcer(authModel, authPolicy)
	if err != nil {
		fmt.Println("Error when building enforcer:", err)
		return nil, err
	}
	return enforcer, nil
}

func sessionIsValidWith(authToken string) (bool, errors.Error) {
	if claims, err := utils.ExtractTokenClaims(authToken); err != nil {
		return false, nil
	} else if uID, err := uuid.Parse(claims.AccountID); err != nil {
		return false, nil
	} else if exists, err := authService.SessionExists(uID, authToken); err != nil {
		return false, err
	} else if !exists {
		return false, nil
	}
	return true, nil
}
