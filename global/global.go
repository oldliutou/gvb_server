package global

import (
	"gvb_server/config"

	"gorm.io/gorm"
)

/*
	全局变量文件
*/
var (
	Config *config.Config // 配置文件
	DB *gorm.DB           // 数据库
)