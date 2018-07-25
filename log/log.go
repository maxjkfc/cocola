package log

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
)

var log Logger

type Logger interface {
	Debug(msg string, output interface{})
	Info(msg string, output interface{})
	Warn(msg string, output interface{})
	Error(msg string, output interface{})
	Panic(msg string, output interface{})
	Fatal(msg string, output interface{})
}

func init() {
	defaultLog()
}

func Log() Logger {
	return log
}

func defaultLog() {
	newZapConfig().Build()
}
