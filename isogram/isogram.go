package isogram

import "strings"

// IsIsogram checks if all characters in a word are contained only once
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for _, character := range word {
		characterString := string(character)

		if strings.ContainsAny(characterString, "- ") {
			continue
		}
		counts := strings.Count(word, characterString)
		if counts > 1 {
			return false
		}
	}
	return true
}
