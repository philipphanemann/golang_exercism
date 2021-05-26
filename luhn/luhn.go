package luhn

import (
	"strings"
	"unicode"
)

func ReverseString(s string) string {
	N := len(s)
	reversed := make([]rune, N)

	for i, r := range s {
		reversed[N-i-1] = r
	}
	return string(reversed)
}

func CalcValue(r rune, i int) int {
	val := int(r - '0')
	if i%2 == 0 {
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
	if len(s) <= 1 {
		return false
	}
	reversed := ReverseString(s)

	var sum int
	for i, r := range reversed {
		if !unicode.IsDigit(r) {
			return false
		}

		sum += CalcValue(r, i)
	}

	if sum%10 == 0 {
		return true
	}
	return false

}
