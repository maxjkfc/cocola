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

func Debug(msg string, output interface{}) {
	log.Debug(msg, output)
}

func Info(msg string, output interface{}) {
	log.Info(msg, output)
}

func Warn(msg string, output interface{}) {
	log.Warn(msg, output)
}

func Error(msg string, output interface{}) {
	log.Error(msg, output)
}

func Panic(msg string, output interface{}) {
	log.Panic(msg, output)
}

func Fatal(msg string, output interface{}) {
	log.Fatal(msg, output)
}
