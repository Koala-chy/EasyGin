package main

import (
	"easy-gin/bootstrap"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 实例化 Gin
	r := gin.New()

	// 初始化路由绑定
	bootstrap.InitRoute(r)
	//模版以及静态资源配置
	bootstrap.SetUpTemplate(r)

	// 运行服务，默认为 8080
	err := r.Run()
	if err != nil {
		//打印错误
		fmt.Println(err.Error())
	}

}
