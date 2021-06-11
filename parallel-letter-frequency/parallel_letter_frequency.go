package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(txts []string) FreqMap {
	m := FreqMap{}
	for _, txt := range txts {
		for i, count := range Frequency(txt) {
			m[i] += count
		}
	}

	return m
}
