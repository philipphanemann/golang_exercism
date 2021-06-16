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
	c := make(chan FreqMap, 6)

	for _, txt := range txts {
		go FrequencyGoroutine(txt, c)
	}
	x, y, z := <-c, <-c, <-c
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
