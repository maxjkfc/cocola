package errors

import (
	"fmt"
	"testing"
)

func Test_ErrorNew(t *testing.T) {
	err := New(1, "the test error")

	fmt.Printf("Type: %T\n", err)
	fmt.Println("Error Msg:", err.Error())
	fmt.Println(err.Json())
}

func Test_ErrorList(t *testing.T) {
	t.Log(List())
}
