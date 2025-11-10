package main

import (
	"fmt"
	"go-demo3/internal/config"
	"go-demo3/internal/global"
	"go-demo3/internal/models"
	"go-demo3/internal/router"
	"go-demo3/pkg/db"
	"go-demo3/pkg/logger"

	"go.uber.org/zap"
)

// 程序入口，加载配置、初始化依赖并启动服务
func main() {
	cfg, err := config.LoadConfig("../etc/config.yaml")
	if err != nil {
		panic(fmt.Sprintf("配置加载失败: %v", err))
	}
	global.Config = cfg

	global.Log = logger.NewLogger(&cfg.Log)
	global.LogS = global.Log.Sugar()
	global.Log.Info("配置加载成功", zap.Any("config", global.Config))

	dbIns, err := db.InitDB(&cfg.DB)
	if err != nil {
		global.Log.Error("数据库连接失败", zap.Error(err))
		panic("数据库连接失败")
	}
	global.DB = dbIns
	global.Log.Info("数据库初始化成功")

	err = global.DB.AutoMigrate(&models.Book{})
	if err != nil {
		global.Log.Error("Book表自动迁移失败", zap.Error(err))
		panic("Book表自动迁移失败")
	}
	global.Log.Info("Book表结构同步完成")

	r := router.InitRouter()
	global.Log.Info("服务启动，监听:8080")
	r.Run(":8080")
}
