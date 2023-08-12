package main

import (
	"flag"
	"fmt"

	"easy-gin/bootstrap"
	baseConfig "easy-gin/config"
	"easy-gin/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	//加载 config 目录下的配置信息
	baseConfig.Initialize()
}

func main() {
	//配置初始化
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件")
	flag.Parse()
	config.InitConfig(env)

	// 实例化 Gin
	r := gin.New()

	// 初始化路由绑定
	bootstrap.InitRoute(r)
	//模版以及静态资源配置
	bootstrap.SetUpTemplate(r)

	// 运行服务，默认为 8080
	err := r.Run(":" + config.GetConfig[string]("app.port"))
	if err != nil {
		//打印错误
		fmt.Println(err.Error())
	}

}
