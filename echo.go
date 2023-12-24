package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))


}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
