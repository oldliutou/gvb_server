package config

import "strconv"

type Mysql struct {
	Host string `yaml: "host"`
  	Port int `yaml: "port"`
	Config string `yaml: "config"` // 高级配置，例如 charset
	DB string `yaml: "db"`
	User string `yaml: "db"`
	Password string `yaml: "password"`
	LogLevel string `yaml: "loglevel"` // 日志等级
}

func (m *Mysql) Dsn() string {
    
	return  m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port)  + ")/" + m.DB + "?" + m.Config

}