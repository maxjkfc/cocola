package currency

import "testing"

func Test_IntToString(t *testing.T) {
	x := 1000000

	t.Log(IntToString(x))
	t.Log(IntToString(24885020938402935))
	t.Log(IntToString(10000))
}
