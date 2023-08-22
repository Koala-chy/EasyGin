// Package routes 注册 api 路由
package routes

import (
	"easy-gin/app/controllers/api/auth"
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

		//auth 接口组
		authGroup := apiGroup.Group("/auth")
		{
			rc := new(auth.RegisterController)
			authGroup.POST("/register/phone/exist", rc.CheckPhoneExist)
		}

	}
}
