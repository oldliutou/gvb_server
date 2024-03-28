package config

type Config struct {
    Mysql Mysql `yaml: "mysql"`
	Log Logger `yaml: "log"`
	System System `yaml: "system"`
}





