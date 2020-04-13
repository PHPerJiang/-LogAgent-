package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/gindemo11", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.GET("/gindemo11/a", func(c *gin.Context) {
		c.Request.URL.Path = "/gindemo11/b"
		r.HandleContext(c)
	})
	r.GET("/gindemo11/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, `this is page b`)
	})
	r.Run(":8000")
}
