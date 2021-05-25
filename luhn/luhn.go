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
	} else {
		if val < 5 {
			return 2 * val
		}
	}

	return 7
}

func Valid(s string) bool {
	return true
}
