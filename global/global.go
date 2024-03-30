package global

import (
	"gvb_server/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

/*
	全局变量文件
*/
var (
	Config *config.Config // 配置文件
	DB *gorm.DB           // 数据库
	Log *logrus.Logger    // 日志
)