package clock

import "fmt"

const base int = 60
const minutesPerDay int = 60 * 24

// OwnTime time in minutes
type OwnTime struct {
	m int
}

// New returns positive time of time.
func New(h, m int) OwnTime {
	m = h*base + m
	m %= minutesPerDay
	if m < 0 {
		return OwnTime{m + minutesPerDay}
	}
	return OwnTime{m}
}

// String formats time output
func (t OwnTime) String() string {
	hours := t.m / base
	minutes := t.m % base
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

// Add adds minutes to time
func (t OwnTime) Add(minutes int) OwnTime {
	return New(0, t.m+minutes)
}

// Subtract subtracts minutes from time
func (t OwnTime) Subtract(minutes int) OwnTime {
	return New(0, t.m-minutes)
}
