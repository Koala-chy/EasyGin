// Package bootstrap 模版初始化配置
package bootstrap

import (
	"github.com/gin-gonic/gin"
)

// SetUpTemplate 模版以及静态资源配置
func SetUpTemplate(router *gin.Engine) {
	//静态资源路径配置
	router.Static("/static", "./public/static")
	//模版路径配置
	router.LoadHTMLGlob("templates/**/*")
}
