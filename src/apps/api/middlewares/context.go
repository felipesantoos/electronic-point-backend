package middlewares

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/utils"

	"github.com/labstack/echo/v4"
)

func EnhanceContext(next handlers.RichHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		_, token := utils.ExtractToken(authHeader)
		claims, _ := utils.ExtractTokenClaims(token)
		enhancedCtx, err := handlers.NewRichContext(ctx, claims)
		if err != nil {
			return err
		}
		return next(enhancedCtx)
	}
}
