package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

var counter FreqMap

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func FrequencyGoroutine(s string, c chan FreqMap) {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}

	c <- m
}

func ConcurrentFrequency(txts []string) FreqMap {
	counts := FreqMap{}
	chan_euro := make(chan FreqMap)
	chan_dutch := make(chan FreqMap)
	chan_us := make(chan FreqMap)

	for i, txt := range txts {
		switch i {
		case 0:
			go FrequencyGoroutine(txt, chan_euro)
		case 1:
			go FrequencyGoroutine(txt, chan_dutch)
		case 2:
			go FrequencyGoroutine(txt, chan_us)
		}
	}

	x, y, z := <-chan_euro, <-chan_dutch, <-chan_us
	for i, count := range x {
		counts[i] += count
	}

	for i, count := range y {
		counts[i] += count
	}

	for i, count := range z {
		counts[i] += count
	}
	return counts
}
