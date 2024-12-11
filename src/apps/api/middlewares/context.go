package middlewares

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/utils"
	"eletronic_point/src/core/domain/authorization"

	"github.com/labstack/echo/v4"
)

func EnhanceContext(next handlers.RichHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var claims *authorization.AuthClaims = nil
		if v, ok := ctx.Get("authenticated").(bool); ok && v {
			tokenCookie, _ := ctx.Cookie(handlers.COOKIE_TOKEN_NAME)
			claims, _ = utils.ExtractTokenClaims(tokenCookie.Value)
		}
		enhancedCtx, err := handlers.NewRichContext(ctx, claims)
		if err != nil {
			return err
		}
		return next(enhancedCtx)
	}
}
