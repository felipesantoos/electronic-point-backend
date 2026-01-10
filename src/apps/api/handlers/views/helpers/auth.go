package helpers

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/utils"
	"eletronic_point/src/core/domain/authorization"
	"net/http"
)

// PrepareTokenCookie creates an http.Cookie for the authorization token
func PrepareTokenCookie(auth authorization.Authorization) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = handlers.COOKIE_TOKEN_NAME
	cookie.Value = auth.Token()
	cookie.Path = "/"
	cookie.HttpOnly = false // Allow HTMX to read it if needed, or set to true for security
	cookie.Secure = utils.IsAPIInProdMode()
	cookie.Expires = *auth.ExpirationTime()
	return cookie
}

// PrepareLogoutCookie creates an expired http.Cookie to clear the token
func PrepareLogoutCookie() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = handlers.COOKIE_TOKEN_NAME
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = false
	cookie.MaxAge = -1
	return cookie
}
