package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/out", "./statics")
	router.SetFuncMap(template.FuncMap{
		"showHtml": func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	router.LoadHTMLGlob("templates/*")
	router.GET("/demo8", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.Run(":8000")
}
