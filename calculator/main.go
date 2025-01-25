package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/add/:first/:second", addition)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func addition(c echo.Context) error {
	a, err := strconv.ParseInt(c.Param("first"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "First argument is not a number.")
	}
	b, err := strconv.ParseInt(c.Param("second"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Second argument is not a number.")
	}
	return c.String(http.StatusOK, strconv.FormatInt(int64(a+b), 10))
}
