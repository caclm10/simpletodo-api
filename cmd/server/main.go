package main

import (
	"net/http"

	"github.com/caclm10/simpletodo-api/internal/app"
	"github.com/caclm10/simpletodo-api/internal/config"
	"github.com/labstack/echo/v4"
)

func main() {
	// ----------------------------------------------------
	// Server configuration
	// ----------------------------------------------------
	config.NewViper()
	app.ConnectDB()

	// ----------------------------------------------------
	// Web configuration
	// ----------------------------------------------------
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "Hello World!"})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
