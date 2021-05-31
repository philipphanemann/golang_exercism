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
	var res uint64
	res = uint64(math.Pow(2, 63))
	return 2*res - 1

}
