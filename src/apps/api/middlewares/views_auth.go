package middlewares

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// ViewsAuthorize is a middleware that checks for authentication in cookies
// and redirects to /login if not authenticated. It also handles HTMX requests.
func ViewsAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Check for token in cookie
		cookie, err := ctx.Cookie(handlers.COOKIE_TOKEN_NAME)
		var token string
		if err == nil {
			token = cookie.Value
		}

		// If no token in cookie, check Authorization header (for HTMX requests that might pass it)
		if token == "" {
			authHeader := ctx.Request().Header.Get("Authorization")
			_, token = utils.ExtractToken(authHeader)
		}

		// Validate token
		accRole, ok := utils.ExtractAuthorizationAccountRole(token)

		// Check if it's an anonymous request allowed by some other mechanism
		// For now, we only allow if it's not anonymous or if it's a non-protected path
		// But this middleware is intended to be used on PROTECTED view routes.

		if !ok || accRole == "anonymous" {
			return redirectToLogin(ctx)
		}

		// Verify session in Redis
		if valid, _ := sessionIsValidWith(token); !valid {
			return redirectToLogin(ctx)
		}

		// Only allow ADMIN users for the frontend
		if strings.ToLower(accRole) != "admin" {
			// Clear cookie and redirect to login if not admin
			cookie := &http.Cookie{
				Name:     handlers.COOKIE_TOKEN_NAME,
				Value:    "",
				Path:     "/",
				HttpOnly: true,
				MaxAge:   -1,
			}
			ctx.SetCookie(cookie)
			return redirectToLogin(ctx)
		}

		// Set token in Authorization header for other middlewares/handlers that might expect it
		ctx.Request().Header.Set("Authorization", "Bearer "+token)

		return next(ctx)
	}
}

func redirectToLogin(ctx echo.Context) error {
	// If it's an HTMX request, we can't just redirect with 302
	// because HTMX will follow the redirect and swap the login page into the target area.
	// We want the whole page to redirect.
	if ctx.Request().Header.Get("HX-Request") == "true" {
		ctx.Response().Header().Set("HX-Redirect", "/login")
		return ctx.NoContent(http.StatusUnauthorized)
	}

	return ctx.Redirect(http.StatusFound, "/login")
}

// AdminAuthorize is a middleware that checks if the user has the ADMIN role
func AdminAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// This should be used AFTER ViewsAuthorize to ensure token is present
		authHeader := ctx.Request().Header.Get("Authorization")
		_, token := utils.ExtractToken(authHeader)

		accRole, _ := utils.ExtractAuthorizationAccountRole(token)
		if strings.ToLower(accRole) != "admin" {
			return ctx.NoContent(http.StatusForbidden)
		}

		return next(ctx)
	}
}
