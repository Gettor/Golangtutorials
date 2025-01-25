package main

import (
	"fmt"
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
	e.GET("/sub/:first/:second", subtraction)
	e.GET("/mult/:first/:second", multiplication)
	e.GET("/div/:first/:second", division)

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

func subtraction(c echo.Context) error {
	a, err := strconv.ParseInt(c.Param("first"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "First argument is not a number.")
	}
	b, err := strconv.ParseInt(c.Param("second"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Second argument is not a number.")
	}
	return c.String(http.StatusOK, strconv.FormatInt(int64(a-b), 10))
}

func multiply(left, right int64) (int64, error) {
	const mostPositive = 1<<63 - 1
	const mostNegative = -(mostPositive + 1)
	result := left * right
	if left == 0 || right == 0 || left == 1 || right == 1 {
		return result, nil
	}
	if left == mostNegative || right == mostNegative {
		return result, fmt.Errorf("overflow multiplying %v and %v", left, right)
	}
	if result/right != left {
		return result, fmt.Errorf("overflow multiplying %v and %v", left, right)
	}
	return result, nil
}

func multiplication(c echo.Context) error {
	a, err := strconv.ParseInt(c.Param("first"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "First argument is not a number.")
	}
	b, err := strconv.ParseInt(c.Param("second"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Second argument is not a number.")
	}
	result, err := multiply(a, b)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, strconv.FormatInt(int64(result), 10))
}

func division(c echo.Context) error {
	a, err := strconv.ParseInt(c.Param("first"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "First argument is not a number.")
	}
	b, err := strconv.ParseInt(c.Param("second"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Second argument is not a number.")
	}
	if b == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Division by zero detected")
	}
	return c.String(http.StatusOK, strconv.FormatFloat(float64(float64(a)/float64(b)), 'f', 2, 64))
}
