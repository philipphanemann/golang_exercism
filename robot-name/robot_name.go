package robotname

import (
	"math/rand"
	"time"
)

// Robot type
type Robot struct {
	name string
}

// Name creates a robot name
func (r *Robot) Name() (string, error) {
	rand.Seed(time.Now().UnixNano())
	if r.name == "" {
		r.name = string([]rune{R('A', 25), R('A', 25), R('0', 9), R('0', 9), R('0', 9)})
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
