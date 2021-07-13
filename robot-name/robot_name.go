package robotname

import (
	"math/rand"
	"time"
)

var robots = make(map[*Robot]string)

// Robot type
type Robot struct {
}

//randInt produces a random number within the interval [min, max)
//i.e. to represent the range for AB max should be choosen to be 67
func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// Name creates a robot name
func (r *Robot) Name() (string, error) {
	if val, ok := robots[r]; ok {
		return val, nil
	} else {
		s := string([]byte{
			byte(randInt(65, 91)),
			byte(randInt(65, 91)),
			byte(randInt(48, 58)),
			byte(randInt(48, 58)),
			byte(randInt(48, 58)),
		})

		robots[r] = s
		return s, nil
	}
}

// Reset currently placeholder to compile
func (r *Robot) Reset() {}
