package log

import (
	"testing"

	"go.uber.org/zap"
)

type output struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func Test_NewZapConfig(t *testing.T) {
	if err := NewZapConfig().Env(DEV).Level(DEBUG).Build(); err != nil {
		t.Error(err)
	}
	t.Log("Success")
}

func Test_Zlog(t *testing.T) {
	Zlog().Info("Test", zap.Any("context", 123))
}

func Test_Zlogs(t *testing.T) {
	Zlogs().Error("Test", "test2")
}
