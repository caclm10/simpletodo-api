package helper

import (
	"net/http"

	"github.com/caclm10/simpletodo-api/internal/model/response"
	"github.com/labstack/echo/v4"
)

func BindRequest[V comparable](c echo.Context) (V, error) {
	var b V

	if err := c.Bind(&b); err != nil {
		return b, response.NewHTTPErrorMessageResponse(http.StatusBadRequest, err, "bad request")
	}

	if err := c.Validate(&b); err != nil {
		return b, err
	}

	return b, nil
}
