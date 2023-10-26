package response

import (
	"github.com/caclm10/simpletodo-api/internal/app"
	"github.com/labstack/echo/v4"
)

type DefaultResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type MessageResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func getStatusByCode(code int) string {

	if code >= 200 && code < 300 {
		return "Success"
	} else if code >= 300 && code < 400 {
		return "Redirected"
	}

	return "Error"
}

func NewDefaultResponse(code int, a any) DefaultResponse {
	return DefaultResponse{
		Code:   code,
		Status: getStatusByCode(code),
		Data:   a,
	}
}

func NewMessageResponse(code int, msg string) MessageResponse {
	return MessageResponse{
		Code:    code,
		Status:  getStatusByCode(code),
		Message: msg,
	}
}

func NewHTTPErrorMessageResponse(code int, err error, msg string) error {
	message := msg

	if app.IsDev() {
		message = err.Error()
	}

	return echo.NewHTTPError(code, NewMessageResponse(code, message))
}
