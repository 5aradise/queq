package main

import (
	"errors"
	"math"
)

var errNoRealSolutions = errors.New("equation has no real solutions")

func solve(a, b, c float64) (x1, x2 float64, err error) {
	D := b*b - 4*a*c

	if D < 0 {
		err = errNoRealSolutions
		return
	}

	if D == 0 {
		x1 = -b / (2 * a)
		x2 = x1
		return
	}

	srD := math.Sqrt(D)
	x1 = (-b + srD) / (2 * a)
	x2 = (-b - srD) / (2 * a)
	return
}
