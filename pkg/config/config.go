// Package config 负责调用配置信息
package config

import (
	"fmt"
	"os"

	"easy-gin/pkg/helpers"

	viperLib "github.com/spf13/viper"
)

// 实例化 viper 库
var viper *viperLib.Viper

// ConfigData  动态加载配置信息
type ConfigData func() map[string]interface{}

// ConfigDatas  先加载到此数组 再动态生成配置信息
var ConfigDatas map[string]ConfigData

// 初始化
func init() {
	//初始化 viper 库
	viper = viperLib.New()
	//配置类型 支持类型: json、toml、yaml、yml、properties、props、prop、env、dotenv
	viper.SetConfigType("env")
	//环境变量配置文件查找路径,相对于 main.go
	viper.AddConfigPath(".")
	//设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")
	//读取环境变量
	viper.AutomaticEnv()

	ConfigDatas = make(map[string]ConfigData)
}

// InitConfig 初始化 配置信息 完成对环境变量以及 config 配置信息的加载
func InitConfig(env string) {
	//加载环境变量
	loadEnv(env)
	//注册配置信息
	loadConfig()
}

// 注册配置信息
func loadConfig() {
	for name, con := range ConfigDatas {
		viper.Set(name, con())
	}
}

// 加载 env 配置信息
func loadEnv(envSuffix string) {
	//默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := ".env." + envSuffix
		//Stat 获取文件属性
		if _, err := os.Stat(filePath); err == nil {
			//如 .env.xxx 类型
			envPath = filePath
		}
	}
	//加载 env
	viper.SetConfigName(envPath)
	//ReadInConfig用于读取查找到的配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err)) // 返回失败信息，可能由于文件不存在等
	}
	//监控 .env 文件,变更时重新加载
	viper.WatchConfig()
}

// Env 读取环境变量,支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// AddConfig 新增配置项
func AddConfig(name string, configData ConfigData) {
	ConfigDatas[name] = configData
}

// GetConfig 使用泛型 获取配置项
/**
@params path 允许使用 . 获取配置,如: app.name
@params defaultValue 允许传参默认值
*/
func GetConfig[T any](path string, defaultValue ...interface{}) T {
	if value := internalGet(path, defaultValue...); value != nil {
		return value.(T)
	}
	// 泛型不能返回 nil，因此需要根据类型建立空变量，这样返回的会是对应类型的"空"值
	var fallback T
	return fallback
}

// 获取配置信息
func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

//// GetString 获取 String 类型的配置信息
//func GetString(path string, defaultValue ...interface{}) string {
//	return cast.ToString(internalGet(path, defaultValue...))
//}
//
//// GetInt 获取 Int 类型的配置信息
//func GetInt(path string, defaultValue ...interface{}) int {
//	return cast.ToInt(internalGet(path, defaultValue...))
//}
//
//// GetFloat64 获取 float64 类型的配置信息
//func GetFloat64(path string, defaultValue ...interface{}) float64 {
//	return cast.ToFloat64(internalGet(path, defaultValue...))
//}
//
//// GetInt64 获取 Int64 类型的配置信息
//func GetInt64(path string, defaultValue ...interface{}) int64 {
//	return cast.ToInt64(internalGet(path, defaultValue...))
//}
//
//// GetUint 获取 Uint 类型的配置信息
//func GetUint(path string, defaultValue ...interface{}) uint {
//	return cast.ToUint(internalGet(path, defaultValue...))
//}
//
//// GetBool 获取 Bool 类型的配置信息
//func GetBool(path string, defaultValue ...interface{}) bool {
//	return cast.ToBool(internalGet(path, defaultValue...))
//}
//
//// GetStringMapString 获取结构数据
//func GetStringMapString(path string) map[string]string {
//	return viper.GetStringMapString(path)
//}
