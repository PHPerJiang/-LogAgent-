package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	log.Println("m1 in..")
	c.Set("key", "Gopher")
	log.Println("m1 out ..")
}

func m2(c *gin.Context) {
	log.Println("m2 in ..")
	value, ok := c.Get("key")
	if ok {
		log.Printf("get key in m2 %#v", value)
	}
	c.Set("key", "m2")
	m3(c)
	log.Println("m2 out..")
}

func m3(c *gin.Context) {
	log.Println("m3 in..")
	log.Println("m3 out ..")
}

func main() {
	r := gin.Default()
	r.Use(m1, m2)
	r.GET("/gindemo12", func(c *gin.Context) {
		value, ok := c.Get("key")
		if ok {
			c.JSON(http.StatusOK, value)
		} else {
			c.JSON(http.StatusOK, `no data`)
		}
	})
	r.Run(":8000")
}
