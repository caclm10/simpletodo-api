package action

import (
	"net/http"

	"github.com/caclm10/simpletodo-api/internal/app"
	"github.com/caclm10/simpletodo-api/internal/model"
	"github.com/caclm10/simpletodo-api/internal/model/response"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context, u model.User) error {
	sess, err := session.Get("auth", c)
	if err != nil {
		return response.NewHTTPErrorMessageResponse(http.StatusInternalServerError, err, "internal server error")
	}

	sess.Options = &sessions.Options{
		MaxAge:   86400 * 7,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   app.IsProd(),
		Domain:   app.Config.GetString("SESSION_DOMAIN"),
	}

	sess.Values["user"] = u.ID
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return response.NewHTTPErrorMessageResponse(http.StatusInternalServerError, err, "internal server error")
	}

	return nil
}

func Logout(c echo.Context) error {
	sess, err := session.Get("auth", c)
	if err != nil {
		return response.NewHTTPErrorMessageResponse(http.StatusInternalServerError, err, "internal server error")
	}

	sess.Options.MaxAge = -1
	sess.Values["user"] = nil

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return response.NewHTTPErrorMessageResponse(http.StatusInternalServerError, err, "internal server error")
	}

	return nil
}
