package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var created = map[string]int{}
var ubNames = 26 * 26 * 10 * 10 * 10

// Robot type
type Robot struct {
	name string
}

// Name creates a robot name
func (r *Robot) Name() (string, error) {
	rand.Seed(time.Now().UnixNano())
	var name string
	if r.name == "" {
		if len(created) >= ubNames {
			return "", fmt.Errorf("names are exhausted")
		}
		for {
			name = createName()
			if _, ok := created[name]; ok {
				name = createName()
			} else {
				created[name] = 1
				break
			}

		}
		r.name = name
	}
	return r.name, nil
}

// Reset currently placeholder to compile
func (r *Robot) Reset() {
	r.name = ""
}

// R returns a random rune in range r starting at offset o.
func R(o, r int) rune {
	return rune(rand.Intn(r) + o)
}

//createName randomly creates a name
func createName() string {
	return string([]rune{R('A', 25), R('A', 25), R('0', 9), R('0', 9), R('0', 9)})
}
