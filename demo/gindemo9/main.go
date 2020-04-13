package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Name string
	age  int
}

func main() {
	r := gin.Default()
	r.GET("/gindemo9", func(c *gin.Context) {
		data1 := data{Name: "Gopher", age: 23}
		// c.JSON(http.StatusOK, map[string]interface{}{
		// 	"name": "Gopher",
		// })
		c.JSON(http.StatusOK, data1)
	})
	r.Run(":8000")
}
