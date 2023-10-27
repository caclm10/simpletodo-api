package middleware

import (
	"github.com/caclm10/simpletodo-api/internal/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{app.Config.GetString("FRONTEND_URL")},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PATCH, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentType, echo.HeaderCookie, echo.HeaderOrigin, "X-XSRF-TOKEN"},
		AllowCredentials: true,
	})
}
