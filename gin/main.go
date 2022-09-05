package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/random", answerRandom)
	r.POST("/random", errorPage)

	r.Run(":8084") // listen and serve on 0.0.0.0:8084
}

func answerRandom(c *gin.Context) {
	// Get team and member from the query string

	smin := c.Query("min")
	smax := c.Query("max")

	if smin == "" || smax == "" {
		c.String(404, "bad request: params are not diffined")
		return
	}
	min, err := strconv.Atoi(smin)
	if err != nil {
		c.String(404, "impossible convert to integer")
		return
	}
	max, err := strconv.Atoi(smax)
	if err != nil {
		c.String(404, "impossible convert to integer")
		return
	}
	if min >= max {
		c.String(404, "bad request: params are not diffined")
		return
	}
	result := strconv.Itoa(rand.Intn(max-min+1) + min)

	c.String(http.StatusOK, "result: "+result)
}

func errorPage(c *gin.Context) {
	c.String(http.StatusBadRequest, "Bad method or bad request")
}
