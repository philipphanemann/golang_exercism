package grains

//package main

import (
	"errors"
	"math"
)

func Square(x int) (square uint64, err error) {
	if x < 1 || x > 64 {
		return 0, errors.New("input is out of range")
	}
	return uint64(math.Pow(2, float64(x-1))), nil
}

func Total() uint64 {
	var sum uint64
	for i := 1; i < 65; i++ {
		calc, _ := Square(i)
		sum += calc
	}
	return sum

}
