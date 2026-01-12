package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCORSMiddleware(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderOrigin, "http://localhost")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := CORSMiddleware()(func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	err := handler(c)
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost", rec.Header().Get(echo.HeaderAccessControlAllowOrigin))
}
