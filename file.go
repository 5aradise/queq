package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type fileMode struct {
	in  bufio.Scanner
	out io.Writer
}

func newFileMode(input io.Reader, output io.Writer) mode {
	return &fileMode{*bufio.NewScanner(input), output}
}

func (m *fileMode) run() error {
	for {
		a, b, c, err := m.readLine()
		if err != nil {
			return fmt.Errorf("invalid file format: %w", err)
		}

		_, err = m.out.Write(fmtEquation(a, b, c))
		if err != nil {
			return err
		}

		x1, x2, ok, err := solve(a, b, c)
		if err != nil {
			return err
		}

		_, err = m.out.Write(append(fmtSolution(x1, x2, ok), '\n'))
		if err != nil {
			return err
		}
	}
}

func (m *fileMode) readLine() (a, b, c float64, err error) {
	m.in.Scan()
	line := m.in.Text()
	vars := strings.Fields(line)
	if len(vars) != 3 {
		err = fmt.Errorf("invalid number of parameters on line: %s", line)
		return
	}

	a, err = strconv.ParseFloat(vars[0], 64)
	if err != nil {
		return
	}
	b, err = strconv.ParseFloat(vars[1], 64)
	if err != nil {
		return
	}
	c, err = strconv.ParseFloat(vars[2], 64)
	if err != nil {
		return
	}

	return
}
