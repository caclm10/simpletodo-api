package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/caclm10/simpletodo-api/internal/app"
	"github.com/caclm10/simpletodo-api/internal/middleware"
	"github.com/caclm10/simpletodo-api/internal/model"
	"github.com/caclm10/simpletodo-api/internal/model/response"
	"github.com/caclm10/simpletodo-api/internal/service"
	"github.com/caclm10/simpletodo-api/internal/validation"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSignUpShouldSuccess(t *testing.T) {
	app.NewConfig("..", "..")
	app.ConnectTestDB()
	e := echo.New()
	e.Validator = validation.NewValidator()
	e.Use(middleware.Session())

	s := service.NewAuthService(app.DB)
	h := NewAuthHandler(s)

	e.POST("/", h.SignUp)

	req, _ := http.NewRequest(echo.POST, "/", strings.NewReader(`{ "name": "Nanashi Mumei", "email": "nanashimumei@gmail.com", "password": "mumei" }`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	res := httptest.NewRecorder()

	e.ServeHTTP(res, req)

	if assert.Equal(t, http.StatusCreated, res.Code) {
		var resp response.DefaultResponse
		json.Unmarshal(res.Body.Bytes(), &resp)

		userRes := resp.Data.(map[string]interface{})["user"].(map[string]interface{})
		assert.Equal(t, "nanashimumei@gmail.com", userRes["email"])

		user := model.User{ID: uint(userRes["id"].(float64))}
		assert.NoError(t, app.DB.First(&user).Error)
	}

	app.DBRenewTable("users")

	e.Close()
}
