package global

import (
	"go-demo3/internal/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局配置变量
var Config *config.Config

// 全局日志
var Log *zap.Logger

// 全局sugar日志
var LogS *zap.SugaredLogger

// 全局数据库DB变量
var DB *gorm.DB
