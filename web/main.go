package main

import (
	"net/http"

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

type Persion struct {
	Name string `json:"name" binding:"required" uri:"name"`
	Age  int    `json:"age" uri:"age"`
}

//BindJsonPostTest 绑定参数的请求获取
func BindJsonPostTest() {
	var p Persion
	router.POST("/", func(c *gin.Context) {
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(200, p)
	})
}

func BindUrlTest() {
	var p Persion
	router.GET("/:name/:age", func(c *gin.Context) {
		if err := c.ShouldBindUri(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(200, p)
	})
}

func main() {
	Init()
	// GetTest()
	// DefaultGetTest()
	// PostGroup()
	// BindJsonPostTest()
	BindUrlTest()
	Run()
}
