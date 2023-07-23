// Package bootstrap 处理程序的初始化逻辑
package bootstrap

import (
	"easy-gin/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// InitRoute 初始化路由
func InitRoute(router *gin.Engine) {
	//注册全局中间件
	registerGlobalMiddleWare(router)
	//  注册 API 路由
	routes.RegisterApiRoutes(router)
	//  配置 404 路由
	setup404Handler(router)
}

// 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

// 处理 404 请求
func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			//如果 Html 的话
			c.HTML(http.StatusNotFound, "error/404.html", gin.H{
				"title": "EasyGin",
			})
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "NotFound 路由不存在",
			})
		}
	})
}
