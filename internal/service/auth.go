package service

import (
	"net/http"

	"github.com/caclm10/simpletodo-api/internal/action"
	"github.com/caclm10/simpletodo-api/internal/lib"
	"github.com/caclm10/simpletodo-api/internal/model"
	"github.com/caclm10/simpletodo-api/internal/model/query"
	"github.com/caclm10/simpletodo-api/internal/model/request"
	"github.com/caclm10/simpletodo-api/internal/model/response"
	"github.com/caclm10/simpletodo-api/internal/validation"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db}
}

func (s *AuthService) SignUp(c echo.Context, b request.SignUpRequest) (response.UserResponse, error) {
	ur := response.UserResponse{}

	u, found, err := query.FindUserByEmail(s.db, b.Email)
	if err != nil {
		return ur, response.NewHTTPErrorMessageResponse(http.StatusInternalServerError, err, "internal server error")
	}

	if found {
		return ur, response.NewHTTPValidationErrorResponse(validation.ValidationErrors{{
			Field:   "email",
			Message: "Email already in used.",
		}})
	}

	password, err := lib.MakeHash(b.Password)
	if err != nil {
		return ur, response.NewHTTPErrorMessageResponse(http.StatusInternalServerError, err, "internal server error")
	}

	u = model.User{
		Name:     b.Name,
		Email:    b.Email,
		Password: password,
	}

	if err := s.db.Create(&u).Error; err != nil {
		return ur, response.NewHTTPErrorMessageResponse(http.StatusInternalServerError, err, "internal server error")
	}

	if err := action.Login(c, u); err != nil {
		return ur, err
	}

	return u.ToResponse(), nil
}

func (s *AuthService) SignIn(c echo.Context, b request.SignUpRequest) (response.UserResponse, error) {
	ur := response.UserResponse{}

	u, found, err := query.FindUserByEmail(s.db, b.Email)
	if err != nil {
		return ur, response.NewHTTPErrorMessageResponse(http.StatusInternalServerError, err, "internal server error")
	}

	if !found {
		return ur, response.NewHTTPValidationErrorResponse(validation.ValidationErrors{{
			Field:   "email",
			Message: "Email not registered.",
		}})
	}

	if !lib.CheckHash(b.Password, u.Password) {
		return ur, response.NewHTTPValidationErrorResponse(validation.ValidationErrors{{
			Field:   "password",
			Message: "Incorrect password.",
		}})
	}

	if err := action.Login(c, u); err != nil {
		return ur, err
	}

	return u.ToResponse(), nil
}

func (s *AuthService) SignOut(c echo.Context) error {
	if err := action.Logout(c); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) User(c echo.Context) response.UserResponse {
	u := c.Get("user").(model.User)

	return u.ToResponse()
}
