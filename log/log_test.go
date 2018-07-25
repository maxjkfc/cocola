package log

import (
	"testing"
)

type output struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func Test_initLog(t *testing.T) {
	Log().Info("inti", nil)

}

func Test_NewZapConfig(t *testing.T) {
	if _, err := NewZapConfig().Env(DEV).Level(DEBUG).Build(); err != nil {
		t.Error(err)
	} else {

		t.Log("Success")

	}
}

func Test_Zlog_Output(t *testing.T) {

	Log().Info("Bug", output{
		Code: "1",
		Msg:  "kkuu",
	})

}

func Test_Zlog(t *testing.T) {
	x := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		"jimmy",
		10,
	}
	Log().Info("Test", x)
}
