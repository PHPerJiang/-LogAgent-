package main

import (
	"log"
	"net/http"
	"time"

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

// LoadHtmlGlob 加载模板文件夹
func LoadHtmlGlob() {
	router.LoadHTMLGlob("templates/*")
	router.GET("/web", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "i am Gopher!",
		})
	})
}

//LoadHtmlFiles 加载模板文件
func LoadHtmlFiles() {
	router.LoadHTMLFiles("templates/index.html")
	router.GET("/web", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "i am Gopher!",
		})
	})
}

// Redirect 重定向
func Redirect() {
	router.GET("/web", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
}

//AsyncGoroutines  异步请求
func AsyncGoroutines() {
	router.GET("/async", func(c *gin.Context) {
		cp := c.Copy() //异步必须用content的拷贝值
		go func() {
			time.Sleep(time.Second * 3)
			log.Println(cp.Request.URL.Path)
		}()
	})
}

//SyncGoroutines 同步请求
func SyncGoroutines() {
	router.GET("/sync", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
}

//MiddleWare 中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		c.Set("middleName", "MiddleWare")
		c.Next() //自定义中间件必须要使用next函数加载
		t := time.Since(now)
		log.Printf("load middleName speed:%v", t)
	}
}

//UseMiddleWare 使用中间件
func UseMiddleWare() {
	router.Use(MiddleWare())
	router.GET("/middleware", func(c *gin.Context) {
		name, _ := c.Get("middleName")
		c.JSON(http.StatusOK, gin.H{"middleName:": name})
	})
}

//Cookie 使用cookie
func Cookie() {
	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("Gopher")
		if err != nil {
			log.Printf("no cookie %v", err)
			c.SetCookie("Gopher", "value", 60, "/", "", false, true)
			return
		}
		log.Printf("cookie: %v", cookie)
	})
}

//CookieMiddleWare cookie中间件
func CookieMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("Gopher", "Gopher", 60, "/", "", false, true)
		name, err := c.Cookie("Gopher")
		if err != nil {
			log.Printf("no cookie! err:%v", err)
			c.JSON(http.StatusNonAuthoritativeInfo, "failed")
			c.Abort()
		}
		if name != "Gopher" {
			c.JSON(http.StatusNonAuthoritativeInfo, "failed")
			c.Abort()
		}
		c.Next()
	}
}

//UseCookieMiddleWare 使用cookie中间件
func UseCookieMiddleWare() {
	router.GET("/cookiemiddleware", CookieMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})
}

func main() {
	Init()
	// GetTest()
	// DefaultGetTest()
	// PostGroup()
	// BindJsonPostTest()
	// BindUrlTest()
	// LoadHtmlGlob()
	// LoadHtmlFiles()
	// AsyncGoroutines()
	// SyncGoroutines()
	// UseMiddleWare()
	// Cookie()
	UseCookieMiddleWare()
	Run()
}
