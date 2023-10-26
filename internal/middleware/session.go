package middleware

import (
	"github.com/caclm10/simpletodo-api/internal/app"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Session() echo.MiddlewareFunc {
	return session.Middleware(sessions.NewCookieStore([]byte(app.Config.GetString("APP_KEY"))))
}
