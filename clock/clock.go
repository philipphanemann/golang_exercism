package clock

import "fmt"

const base int = 60
const MinutesPerDay int = 60 * 24

type OwnTime struct {
	h, m int
}

// MinutesToTime returns positive hour and minutes
func MinutesToTime(m int) (hour int, min int) {
	m %= MinutesPerDay
	if m < 0 {
		m = MinutesPerDay + m
	}
	return m / base, m % base
}

func New(h, m int) OwnTime {
	totalMinutes := h*base + m
	h, m = MinutesToTime(totalMinutes)
	return OwnTime{h, m}
}

func (t OwnTime) String() string {
	return fmt.Sprintf("%02d:%02d", t.h, t.m)
}
