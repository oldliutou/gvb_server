package config

type Logger struct {
	Level        string `yaml: "level"`
	Prefix       string `yaml: "prefix"`
	Director     string `yaml: "director"`
	ShowLine     bool   `yaml: "showline"`     // 是否显示行号
	LogInConsole bool   `yaml: "loginconsole"` //是否显示打印的路径
}
