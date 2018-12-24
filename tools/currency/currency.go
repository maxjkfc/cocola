package currency

import (
	"strconv"
)

func IntToString(f int) string {
	//s := strconv.FormatFloat(f, 'f', 2, 64)
	s := strconv.Itoa(f)
	str := ""
	var x int = 0
	for i := len(s) - 1; i >= 0; i-- {

		if x != 0 && (x%3) == 0 {
			str = "," + str
			x = 0
		}

		x++
		str = string(s[i]) + str
	}
	return str
}
