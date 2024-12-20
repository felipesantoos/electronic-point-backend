package middlewares

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/utils"

	"github.com/labstack/echo/v4"
)

func EnhanceContext(next handlers.RichHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		claims, _ := utils.ExtractTokenClaims(authHeader)
		enhancedCtx, err := handlers.NewRichContext(ctx, claims)
		if err != nil {
			return err
		}
		return next(enhancedCtx)
	}
}
