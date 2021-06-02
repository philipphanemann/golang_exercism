package grains

//package main

import (
	"errors"
)

func Square(x int) (square uint64, err error) {
	if x < 1 || x > 64 {
		return 0, errors.New("input is out of range")
	}
	return uint64(1 << (x - 1)), nil
}

func Total() uint64 {
	return uint64(1<<64 - 1)

}
