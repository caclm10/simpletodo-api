package middleware

import (
	"net/http"

	"github.com/caclm10/simpletodo-api/internal/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CSRF() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "header:X-XSRF-TOKEN",
		CookieName:     "XSRF-TOKEN",
		CookieSameSite: http.SameSiteLaxMode,
		CookieDomain:   app.Config.GetString("SESSION_DOMAIN"),
	})
}
