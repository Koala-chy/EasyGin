// Package config 站点配置信息
package config

import (
	"easy-gin/pkg/config"
)

// 初始化配置信息
func init() {
	//添加配置信息
	config.AddConfig("app", func() map[string]interface{} {
		return map[string]interface{}{
			//应用名称
			"name": config.Env("APP_NAME", "EasyGin"),
			//当前环境信息:  本地:local, Beta环境:beta, 生产环境:production, 测试环境: test
			"env": config.Env("APP_ENV", "production"),
			// 是否进入调试模式
			"debug": config.Env("APP_DEBUG", false),
			// 应用服务端口
			"port": config.Env("APP_PORT", "8080"),
			// 加密会话、JWT 加密
			"key": config.Env("APP_KEY", "cC8cS1qF4eE9gN3eD2oD1bB1jP2pB1dJ"),
			// 用以生成链接
			"url": config.Env("APP_URL", "http://localhost:8080"),
			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
