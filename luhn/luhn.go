package luhn

import (
	"strings"
	"unicode"
)

func CalcValue(r rune, double bool) int {
	val := int(r - '0')
	if !double {
		return val
	}

	if twoVal := 2 * val; twoVal < 10 {
		return twoVal
	} else {
		return twoVal - 9
	}
}

func Valid(s string) bool {

	s = strings.ReplaceAll(s, " ", "")
	N := len(s)
	if N <= 1 {
		return false
	}
	doublePosition := (N%2 == 0)

	var sum int
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}

		sum += CalcValue(r, doublePosition)
		doublePosition = !doublePosition
	}

	return sum%10 == 0

}
