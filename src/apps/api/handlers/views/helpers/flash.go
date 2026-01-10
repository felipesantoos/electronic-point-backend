package helpers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

const FLASH_COOKIE_NAME = "ep_flash"

// SetFlashMessage sets a flash message in a cookie
func SetFlashMessage(c echo.Context, msgType, content string) {
	msg := FlashMessage{Type: msgType, Content: content}
	data, _ := json.Marshal(msg)
	encoded := base64.StdEncoding.EncodeToString(data)

	cookie := &http.Cookie{
		Name:     FLASH_COOKIE_NAME,
		Value:    encoded,
		Path:     "/",
		HttpOnly: false, // Accessible by JS to show toast
	}
	c.SetCookie(cookie)
}

// GetFlashMessages retrieves and clears flash messages from cookies
func GetFlashMessages(c echo.Context) []FlashMessage {
	cookie, err := c.Cookie(FLASH_COOKIE_NAME)
	if err != nil {
		return nil
	}

	// Clear the cookie
	clearCookie := &http.Cookie{
		Name:     FLASH_COOKIE_NAME,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false,
	}
	c.SetCookie(clearCookie)

	decoded, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil
	}

	var msg FlashMessage
	if err := json.Unmarshal(decoded, &msg); err != nil {
		return nil
	}

	return []FlashMessage{msg}
}
