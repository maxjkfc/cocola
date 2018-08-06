package errors

import (
	"fmt"
	"testing"
)

func Test_ErrorNew(t *testing.T) {
	New(1, "the test error")

	err := Err(1)
	fmt.Printf("Type: %T\n", err)
	fmt.Println("Error Msg:", err.Error())
	fmt.Println(err.Json())
}

func Test_ErrorList(t *testing.T) {
	List()
}
