
package config

type Mysql struct {
	Host string `yaml: "host"`
  	Port int `yaml: "port"`
	DB string `yaml: "db"`
	User string `yaml: "db"`
	Password string `yaml: "password"`
	LogLevel string `yaml: "loglevel"` // 日志等级
}