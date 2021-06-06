package clock

import "fmt"

const base int = 60
const minutesPerDay int = 60 * 24

// OwnTime time in minutes
type OwnTime struct {
	m int
}

// normalizeMinutes normalizes to Minutes of a day
func normalizeMinutes(m int) int {
	m %= minutesPerDay
	if m < 0 {
		return m + minutesPerDay
	}
	return m
}

// New returns positive time of time.
func New(h, m int) OwnTime {
	totalMinutes := h*base + m
	return OwnTime{normalizeMinutes(totalMinutes)}
}

// String formats time output
func (t OwnTime) String() string {
	hours := t.m / base
	minutes := t.m % base
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

// Add adds minutes to time
func (t OwnTime) Add(minutes int) OwnTime {
	return OwnTime{normalizeMinutes(t.m + minutes)}
}

// Subtract subtracts minutes from time
func (t OwnTime) Subtract(minutes int) OwnTime {
	return OwnTime{normalizeMinutes(t.m - minutes)}
}
