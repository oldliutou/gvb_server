package core

import (
	"bytes"
	"fmt"
	"gvb_server/global"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// 颜色
const (
	red = 31
	yellow = 33
	blue = 36
	gray = 37
)

type LogFormatter struct {}

func (t *LogFormatter) Format(entry *logrus.Entry)([]byte, error){
	// 根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	    case logrus.DebugLevel,logrus.TraceLevel:
	    	levelColor = gray
		case logrus.WarnLevel:
			levelColor = yellow
		case logrus.ErrorLevel,logrus.FatalLevel,logrus.PanicLevel:
			levelColor = red
		default:
			levelColor = blue
	
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	}else {
		b = &bytes.Buffer{}
	}

	log := global.Config.Logger

	// 设置时间格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
	    // 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b,"%s [%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix,timestamp,levelColor,entry.Level,funcVal,fileVal,entry.Message)
	}else{
		fmt.Fprintf(b,"%s [%s] \x1b[%dm[%s]\x1b[0m %s\n",log.Prefix, timestamp,levelColor,entry.Level,entry.Message)
	}
	return b.Bytes(),nil
}

func InitLogger() *logrus.Logger{
	mLog := logrus.New() //新建一个实例
	mLog.SetOutput(os.Stdout) //设置输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) // 设置自己定义的Formatter
	level,err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		// 如果没有获取到正确的level,则手动设置为info等级
	    level = logrus.InfoLevel
	}
	mLog.SetLevel(level) //设置最低的Level
	InitDefaultLogger() // 初始化默认的log
	return mLog
}
func InitDefaultLogger() {
    // 全局log
	logrus.SetOutput(os.Stdout) //设置输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{}) // 设置自己定义的Formatter
	level,err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		// 如果没有获取到正确的level,则手动设置为info等级
	    level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //设置最低的Level
}