package clock

import "fmt"

const base int = 60
const MinutesPerDay int = 60 * 24

type OwnTime struct {
	h, m int
}

// MinutesToTime returns positive hour and minutes
func MinutesToTime(m int) OwnTime {
	m %= MinutesPerDay
	if m < 0 {
		m = MinutesPerDay + m
	}
	return OwnTime{m / base, m % base}
}

func TimeToMinutes(t OwnTime) int {
	return t.h*base + t.m
}
func New(h, m int) OwnTime {
	totalMinutes := TimeToMinutes(OwnTime{h, m})
	return MinutesToTime(totalMinutes)
}

func (t OwnTime) String() string {
	return fmt.Sprintf("%02d:%02d", t.h, t.m)
}

func (t OwnTime) Add(minutes int) OwnTime {
	m := TimeToMinutes(t) + minutes
	return MinutesToTime(m)
}
