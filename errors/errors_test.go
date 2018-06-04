package errors

import (
	"fmt"
	"testing"
)

func Test_ErrorNew(t *testing.T) {
	New(1, "the test error")
	New(2, "the test error2")
	New(3, "the test error3")

	err := Err(1)
	fmt.Printf("Type: %T\n", err)
	fmt.Println("Error Code:", err.Code())
	fmt.Println("Error Msg:", err.Error())
}

func Test_ErrorList(t *testing.T) {

	List()
}
