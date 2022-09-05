package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/random", answerRandom)
	e.POST("/random", errorPage)
	e.Logger.Fatal(e.Start(":8083"))
}

func answerRandom(c echo.Context) error {
	// Get team and member from the query string

	smin := c.QueryParam("min")
	smax := c.QueryParam("max")

	if smin == "" || smax == "" {
		return c.String(http.StatusBadRequest, "bad request: params are not diffined")
	}
	min, err := strconv.Atoi(smin)
	if err != nil {
		return c.String(http.StatusBadRequest, "impossible convert to integer")
	}
	max, err := strconv.Atoi(smax)
	if err != nil {
		return c.String(http.StatusBadRequest, "impossible convert to integer")
	}
	if min >= max {
		return c.String(http.StatusBadRequest, "bad request: params are not diffined")
	}
	result := strconv.Itoa(rand.Intn(max-min+1) + min)

	return c.String(http.StatusOK, "result: "+result)
}

func errorPage(c echo.Context) error {
	return c.String(http.StatusBadRequest, "Bad method or bad request")
}
