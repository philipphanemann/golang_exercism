// Package twofer greets either a spezified name or 'you'
package twofer

import "fmt"

// ShareWith returns a phrase of a someone calling the function
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
