package middlewares

import (
	"backend_template/src/core/utils"
	"backend_template/src/ui/api/handlers"

	"github.com/labstack/echo/v4"
)

func EnhanceContext(next handlers.EnhancedHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		claims, _ := utils.ExtractTokenClaims(authHeader)
		enhancedCtx, err := handlers.NewEnhancedContext(ctx, claims)
		if err != nil {
			return err
		}
		return next(enhancedCtx)
	}
}
