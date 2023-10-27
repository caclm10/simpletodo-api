package handler

import (
	"net/http"

	"github.com/caclm10/simpletodo-api/internal/helper"
	"github.com/caclm10/simpletodo-api/internal/model/request"
	"github.com/caclm10/simpletodo-api/internal/model/response"
	"github.com/caclm10/simpletodo-api/internal/service"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	b, err := helper.BindRequest[request.SignUpRequest](c)
	if err != nil {
		return err
	}

	u, err := h.service.SignUp(c, b)
	if err != nil {
		return err
	}

	return response.NewHTTPDefaultResponse(c, http.StatusCreated, echo.Map{"user": u})
}

func (h *AuthHandler) SignIn(c echo.Context) error {
	b, err := helper.BindRequest[request.SignInRequest](c)
	if err != nil {
		return err
	}

	u, err := h.service.SignIn(c, b)
	if err != nil {
		return err
	}

	return response.NewHTTPDefaultResponse(c, http.StatusOK, echo.Map{"user": u})
}

func (h *AuthHandler) SignOut(c echo.Context) error {
	if err := h.service.SignOut(c); err != nil {
		return err
	}

	return response.NewHTTPMessageResponse(c, http.StatusOK, "User logged out successfully.")
}

func (h *AuthHandler) User(c echo.Context) error {
	u := h.service.User(c)

	return response.NewHTTPDefaultResponse(c, http.StatusOK, echo.Map{"user": u})
}
