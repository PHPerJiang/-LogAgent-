package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/gindemo10", func(c *gin.Context) {
		name := c.DefaultQuery("name", "")
		c.JSON(http.StatusOK, name)
	})
	r.Run(":8000")
}
