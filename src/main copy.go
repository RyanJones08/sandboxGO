package main

import (
	"fmt"
	"net/http"
	"os"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
