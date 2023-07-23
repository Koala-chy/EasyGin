// Package routes 注册 api 路由
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterApiRoutes 注册 api 相关路由
func RegisterApiRoutes(r *gin.Engine) {
	//注册一个 api 路由组
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello EasyGin",
			})
		})
	}
}
