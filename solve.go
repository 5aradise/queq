package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
)

var errZeroA = errors.New("a cannot be 0")

func solve(a, b, c float64) (x1, x2 float64, ok bool, err error) {
	if a == 0 {
		err = errZeroA
		return
	}

	D := b*b - 4*a*c

	if D < 0 {
		return
	}

	ok = true

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

func fmtEquation(a, b, c float64) []byte {
	return fmt.Appendf(nil, "Equation is: (%s) x^2 + (%s) x + (%s) = 0\n", formatFloat(a), formatFloat(b), formatFloat(c))
}

func fmtSolution(x1, x2 float64, ok bool) []byte {
	b := &bytes.Buffer{}
	if !ok {
		fmt.Fprintln(b, "There are no roots")
		return b.Bytes()
	}

	if x1 == x2 {
		fmt.Fprintln(b, "There is 1 root")
		fmt.Fprintf(b, "x = %s\n", formatFloat(x1))
		return b.Bytes()
	}

	fmt.Fprintln(b, "There are 2 roots")
	fmt.Fprintf(b, "x1 = %s\n", formatFloat(x1))
	fmt.Fprintf(b, "x2 = %s\n", formatFloat(x2))
	return b.Bytes()
}
