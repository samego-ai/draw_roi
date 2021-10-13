package config

import (
	"github.com/kataras/golog"
	"os"
)

type logger struct {
	Path   string
	Level  string
	Prefix string
}

var Logger logger

func init() {
	Logger.Level = os.Getenv("SG_LOGGER_LEVEL")
	Logger.Prefix = os.Getenv("SG_LOGGER_PREFIX")

	// 设置日志级别 && 日志前缀 && 时间格式
	golog.SetLevel(Logger.Level)
	golog.SetPrefix(Logger.Prefix)
	golog.SetTimeFormat("2006-01-02 15:04:05")
}
