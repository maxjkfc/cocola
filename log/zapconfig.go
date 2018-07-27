package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapconfiger struct {
	env      string
	name     string
	level    string
	encoding string
	skip     int
	c        zap.Config
	ec       zapcore.EncoderConfig
	opts     []zap.Option
}

func newZapConfig() Configer {
	return &zapconfiger{
		skip: 1,
		c:    zap.Config{},
		opts: make([]zap.Option, 0, 0),
	}
}

func (zc *zapconfiger) Level(level string) Configer {
	if level == "" {
		zc.level = DEBUG
	} else {
		zc.level = level
	}
	return zc
}

func (zc *zapconfiger) Name(name string) Configer {
	if name == "" {
		zc.name = "zaplog"
	} else {
		zc.name = name
	}
	return zc
}

func (zc *zapconfiger) Env(env string) Configer {
	if env == "" {
		zc.env = DEV
	} else {
		zc.env = env
	}
	return zc
}

func (zc *zapconfiger) AddCallerSkip(skip int) Configer {
	zc.skip = skip
	return zc
}

func (zc *zapconfiger) Build() (Logger, error) {

	switch zc.env {
	case PRD:
		zc.c = productConfig()
	default:
		zc.c = defaultConfig()

		switch zc.level {
		case INFO:
			zc.c.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		case WARN:
			zc.c.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
		case ERROR:
			zc.c.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		case FATAL:
			zc.c.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
		case DEBUG:
			zc.c.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		}
	}

	zc.opts = append(zc.opts, zap.AddCallerSkip(zc.skip))
	return zc.build()
}

func (zc *zapconfiger) build() (Logger, error) {

	x := new(zlog)
	x.zl, x.err = zc.c.Build(zc.opts...)
	log = x
	return x, x.err

}

func defaultEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     rfc3339TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func defaultConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    defaultEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func productConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.WarnLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    defaultEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func rfc3339TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339Nano))
}
