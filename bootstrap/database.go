// Package bootstrap 处理数据库连接操作
package bootstrap

import (
	"errors"
	"fmt"
	"time"

	"easy-gin/pkg/config"
	"easy-gin/pkg/database"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	var dbConfig gorm.Dialector
	switch config.GetConfig[string]("database.connection") {
	//连接数据库类型
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.GetConfig[string]("database.mysql.username"),
			config.GetConfig[string]("database.mysql.password"),
			config.GetConfig[string]("database.mysql.host"),
			config.GetConfig[string]("database.mysql.port"),
			config.GetConfig[string]("database.mysql.database"),
			config.GetConfig[string]("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite
		sqliteDatabase := config.GetConfig[string]("database.sqlite.database")
		dbConfig = sqlite.Open(sqliteDatabase)
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// 设置最大连接数
	database.SqlDB.SetMaxOpenConns(config.GetConfig[int]("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	database.SqlDB.SetMaxIdleConns(config.GetConfig[int]("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SqlDB.SetConnMaxLifetime(time.Duration(config.GetConfig[int]("database.mysql.max_life_seconds")) * time.Second)
}
