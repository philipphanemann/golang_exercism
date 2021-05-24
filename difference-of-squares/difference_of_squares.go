package diffsquares

import "math"

func SquareOfSum(n int) int {
	return int(math.Pow(float64((n * (n + 1) / 2)), 2))
}

func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
