package clock

import "fmt"

const base int = 60
const minutesPerDay int = 60 * 24

// OwnTime time with hours and minutes
type OwnTime struct {
	h, m int
}

// MinutesToTime returns positive hour and minutes
func MinutesToTime(m int) OwnTime {
	m %= minutesPerDay
	if m < 0 {
		m = minutesPerDay + m
	}
	return OwnTime{m / base, m % base}
}

// TimeToMinutes returns total minutes of time
func TimeToMinutes(t OwnTime) int {
	return t.h*base + t.m
}

// New returns positive time of time.
func New(h, m int) OwnTime {
	totalMinutes := TimeToMinutes(OwnTime{h, m})
	return MinutesToTime(totalMinutes)
}

// String formats time output
func (t OwnTime) String() string {
	return fmt.Sprintf("%02d:%02d", t.h, t.m)
}

// Add adds minutes to time
func (t OwnTime) Add(minutes int) OwnTime {
	m := TimeToMinutes(t) + minutes
	return MinutesToTime(m)
}

// Subtract subtracts minutes from time
func (t OwnTime) Subtract(minutes int) OwnTime {
	m := TimeToMinutes(t) - minutes
	return MinutesToTime(m)
}
