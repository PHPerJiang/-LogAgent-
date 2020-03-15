package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//Init 初始化理由
func Init() {
	router = gin.Default()
}

//Run 路由监听端口
func Run() {
	router.Run(":8000")
}

//GetTest get
func GetTest() {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello word")
	})
}

//DefaultGetTest 默认
func DefaultGetTest() {
	router.GET("/defaultGet", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Gopher")
		c.JSON(200, name)
	})
}

//PostGroup 组
func PostGroup() {
	group := router.Group("/v1")
	{
		group.POST("test", func(c *gin.Context) {
			name := c.DefaultPostForm("name", "Gopher")
			c.JSON(200, name)
		})
	}
}

func main() {
	Init()
	// GetTest()
	// DefaultGetTest()
	PostGroup()
	Run()
}
