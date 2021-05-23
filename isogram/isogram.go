package isogram

import "strings"

// IsIsogram checks if all characters in a word are contained only once
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	counter := map[rune]bool{}

	for _, r := range word {
		if r == ' ' || r == '-' {
			continue
		}

		_, err := counter[r]

		if err {
			return false
		}
		counter[r] = true
	}
	return true
}
