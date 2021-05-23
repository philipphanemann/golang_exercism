package isogram

import "strings"

// IsIsogram checks if all characters in a word are contained only once
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	repeated := map[rune]bool{}

	for _, r := range word {
		if r == ' ' || r == '-' {
			continue
		}

		if repeated[r] {
			return false
		}

		repeated[r] = true
	}
	return true
}
