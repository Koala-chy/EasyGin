package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// 实例化 Gin
	r := gin.New()

	//注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	//使用静态资源
	r.Static("/static", "./public/static")
	//使用模版
	r.LoadHTMLGlob("templates/**/*")

	//注册路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "EasyGin",
		})
	})

	//处理 404 请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			//如果 Html 的话
			c.HTML(http.StatusNotFound, "error/404.html", gin.H{
				"title": "EasyGin",
			})
		} else {
			//默认返回 json
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "NotFound 路由不存在",
			})
		}
	})
	// 运行服务，默认为 8080
	r.Run()

}
