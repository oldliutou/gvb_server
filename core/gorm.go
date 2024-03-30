package core

// 连接数据库
import (
	"fmt"
	"gvb_server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		// global.LOG.Warn("mysql配置未填写,取消gorm连接")
	    global.Log.Warnln("mysql配置未填写,取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
	    // 开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	}else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
	    global.Log.Fatalf(fmt.Sprintf("[%s] mysql连接失败",dsn))
	}
	global.Log.Infof(fmt.Sprintf("[%s] mysql连接成功",dsn))
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10) // 设置空闲连接数
	sqlDB.SetMaxOpenConns(100) // 设置最大连接数
	sqlDB.SetConnMaxLifetime(4 * time.Hour) // 设置连接的最大可复用时间

	return db
}

// func MysqlConnect() (*gorm.DB) {
	

// }