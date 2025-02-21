package main

import (
	"errors"
	"fmt"
	"math"
)

var (
	errZeroA           = errors.New("a cannot be 0")
	errNoRealSolutions = errors.New("equation has no real solutions")
)

func solve(a, b, c float64) (x1, x2 float64, err error) {
	if a == 0 {
		err = errZeroA
		return
	}

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
	x1 = (-b - srD) / (2 * a)
	x2 = (-b + srD) / (2 * a)
	return
}

func solveAndPrint(a, b, c float64) error {
	fmt.Printf("Equation is: (%s) x^2 + (%s) x + (%s) = 0\n", formatFloat(a), formatFloat(b), formatFloat(c))
	x1, x2, err := solve(a, b, c)
	if err != nil {
		if !errors.Is(err, errNoRealSolutions) {
			return err
		}

		fmt.Println("There are no roots")
		return nil
	}

	if x1 == x2 {
		fmt.Println("There is 1 root")
		fmt.Printf("x = %s\n", formatFloat(x1))
		return nil
	}

	fmt.Println("There are 2 roots")
	fmt.Printf("x1 = %s\n", formatFloat(x1))
	fmt.Printf("x2 = %s\n", formatFloat(x2))
	return nil
}
