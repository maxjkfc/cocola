package log

import "go.uber.org/zap"

type zlog struct {
	zl  *zap.Logger
	err error
}

func (z *zlog) Debug(msg string, output interface{}) {
	z.zl.Debug(msg, zap.Any("output", output))
}
func (z *zlog) Info(msg string, output interface{}) {
	z.zl.Info(msg, zap.Any("output", output))
}
func (z *zlog) Warn(msg string, output interface{}) {
	z.zl.Warn(msg, zap.Any("output", output))
}
func (z *zlog) Error(msg string, output interface{}) {
	z.zl.Error(msg, zap.Any("output", output))
}
func (z *zlog) Panic(msg string, output interface{}) {
	z.zl.Panic(msg, zap.Any("output", output))
}
func (z *zlog) Fatal(msg string, output interface{}) {
	z.zl.Fatal(msg, zap.Any("output", output))
}
