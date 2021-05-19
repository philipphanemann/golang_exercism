package hamming

import "errors"

// ErrStrandsDiffernet signals error for different strand length
var ErrStrandsDifferent = errors.New("strands have different length")

func Distance(a, b string) (int, error) {
	if a == b {
		return 0, nil
	} else if len(a) != len(b) {
		return 0, ErrStrandsDifferent
	} else {
		lastDistance, err := Distance(a[1:], b[1:])
		if a[0] != b[0] {
			return lastDistance + 1, err
		} else {
			return lastDistance, err
		}
	}
}
