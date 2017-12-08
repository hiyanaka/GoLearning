package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var i int

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		i++
		c.JSON(http.StatusOK, i)
	})
	r.Run()
}
