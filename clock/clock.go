package clock

import "fmt"

// Clock time in minutes
type Clock struct {
	m int
}

// New returns positive time of time.
func New(h, m int) Clock {
	m = h*60 + m
	m %= 60 * 24
	if m < 0 {
		return Clock{m + 60*24}
	}
	return Clock{m}
}

// String formats time output
func (t Clock) String() string {
	hours := t.m / 60
	minutes := t.m % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

// Add adds minutes to time
func (t Clock) Add(minutes int) Clock {
	return New(0, t.m+minutes)
}

// Subtract subtracts minutes from time
func (t Clock) Subtract(minutes int) Clock {
	return New(0, t.m-minutes)
}
