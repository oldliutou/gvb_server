package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	
	// 连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}