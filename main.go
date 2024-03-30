package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"fmt"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}