package response

import (
	"net/http"

	"github.com/caclm10/simpletodo-api/internal/validation"
	"github.com/labstack/echo/v4"
)

type ValidationErrorResponse struct {
	Code   int                         `json:"code"`
	Status string                      `json:"status"`
	Errors validation.ValidationErrors `json:"errors"`
}

func NewValidationErrorResponse(errs validation.ValidationErrors) ValidationErrorResponse {
	return ValidationErrorResponse{
		Code:   http.StatusUnprocessableEntity,
		Status: "Validation error",
		Errors: errs,
	}
}

func NewHTTPValidationErrorResponse(errs validation.ValidationErrors) error {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, NewValidationErrorResponse(errs))
}
