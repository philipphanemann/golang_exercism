package luhn

func ReverseString(s string) string {
	N := len(s)
	reversed := make([]rune, N)

	for i, r := range s {
		reversed[N-i-1] = r
	}
	return string(reversed)
}

func Valid(s string) bool {
	return true
}
