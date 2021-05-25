package luhn

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
	if i%2 == 1 {
		return int(r - '0')
	}

	if twoVal := 2 * val; twoVal < 10 {
		return twoVal
	} else {
		return twoVal - 9
	}
}

func Valid(s string) bool {

	if len(s) <= 1 {
		return false
	}
	reversed := ReverseString(s)

	var sum int
	for i, r := range reversed {
		sum += CalcValue(r, i)
	}

	if sum%10 == 0 {
		return true
	}
	return false

}
