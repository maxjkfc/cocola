package tools

import (
	"strconv"
)

const (
	A = iota + 10
	B
	C
	D
	E
	F
)

var bitCount []uint8

func init() {
	bitCount = []uint8{0, 1, 2, 4, 8}
}

func HexToBits(s string) []uint8 {

}

func BitsToHex(b []uint8) string {
	var s []byte

	var temp uint8

	var j = 0

	for i := len(b); i > 0; i-- {
		if j > 4 {
			j = 1
		}
		temp = b[i] * bitCount[j]
		j++
	}
}

func switchHex(a uint8) string {
	switch a {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		return strconv.Itoa(int(a))
	case A:
		return 'A'
	case B:
		return 'B'
	case C:
		return 'C'
	case D:
		return 'D'
	case E:
	case F:
	}
}
