package log

import "go.uber.org/zap"

/*
1. chosen log type
2. set log config
3. get logger
4. set logger in global
*/

const (
	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
	FATAL = "fatal"
	DEV   = "dev"
	PRD   = "prd"
	ZAP   = "zap"
)

var zlog *zap.Logger

type Configer interface {
	Level(string) Configer
	Name(string) Configer
	Env(string) Configer
	Build() error
}

func NewZapConfig() Configer {
	return newZapConfig()
}

func Zlog() *zap.Logger {
	if zlog == nil {
		newZapConfig().Env(DEV).Build()
	}
	return zlog
}
func Zlogs() *zap.SugaredLogger {
	if zlog == nil {
		newZapConfig().Env(DEV).Build()
	}
	return zlog.Sugar()
}
