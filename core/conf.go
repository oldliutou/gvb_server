package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// 读取yaml文件的配置
func InitConf() {
	// 定义配置文件名称的常量
	// Tips:配置文件的字段不能有-符号，否则会序列化失败
	const ConfigFile = "settings.yaml"
	// 构建配置结构体的实例(实例地址)，后续以便将配置文件中读取数据并反序列化到结构体中，便于后面调用
	c := &config.Config{}
	// 读取配置文件
	yamlConf, err := ioutil.ReadFile(ConfigFile)

	// fmt.Println(string(yamlConf))
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %v", err))
	}
	// 反序列化：将yaml文件中的值反序列化到Config结构体中
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("反序列化配置文件失败: %v", err)
	}
	log.Println("配置yaml文件加载初始化成功")
	fmt.Println(c)
	// 将反序列化后的实例对象赋值给全局配置变量
	global.Config = c
}
