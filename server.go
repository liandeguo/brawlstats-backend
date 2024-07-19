package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Skibidi World")
	})
	e.Logger.Fatal(e.Start(":8123"))
}
