package middlewares

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/utils"
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/role"

	"github.com/labstack/echo/v4"
)

func EnhanceContext(next handlers.RichHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		_, token := utils.ExtractToken(authHeader)

		// If token is missing in header, try cookie (for frontend requests)
		if token == "" {
			if cookie, err := ctx.Cookie(handlers.COOKIE_TOKEN_NAME); err == nil {
				token = cookie.Value
			}
		}

		// Only extract claims if the token is valid
		var richClaims *authorization.AuthClaims
		if token != "" {
			if accRole, ok := utils.ExtractAuthorizationAccountRole(token); ok && accRole != role.ANONYMOUS_ROLE_CODE {
				richClaims, _ = utils.ExtractTokenClaims(token)
			}
		}

		enhancedCtx, err := handlers.NewRichContext(ctx, richClaims)
		if err != nil {
			return err
		}
		return next(enhancedCtx)
	}
}
