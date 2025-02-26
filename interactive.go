package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type interactiveMode struct {
	in  bufio.Scanner
	out io.Writer
}

func newInteractiveMode(input io.Reader, output io.Writer) mode {
	return &interactiveMode{*bufio.NewScanner(input), output}
}

func (m *interactiveMode) run() error {
	for {
		a := m.mustReadFloat("a = ")
		b := m.mustReadFloat("b = ")
		c := m.mustReadFloat("c = ")

		_, err := m.out.Write(fmtEquation(a, b, c))
		if err != nil {
			return err
		}

		x1, x2, ok, err := solve(a, b, c)
		var out []byte
		if err != nil {
			out = append([]byte("Error. "), []byte(err.Error())...)
		} else {
			out = fmtSolution(x1, x2, ok)
		}

		_, err = m.out.Write(append(out, '\n'))
		if err != nil {
			return err
		}
	}
}

func (m *interactiveMode) mustReadFloat(msg string) float64 {
	for {
		fmt.Fprint(m.out, msg)
		m.in.Scan()
		strF := m.in.Text()
		f, err := strconv.ParseFloat(strF, 64)
		if err != nil {
			fmt.Fprintf(m.out, "Error. Expected a valid real number, got %s instead\n", strF)
			continue
		}
		return f
	}
}
