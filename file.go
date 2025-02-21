package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func runFileMode(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	input = bufio.NewScanner(f)

	for input.Scan() {
		line := input.Text()
		a, b, c, err := readLine(line)
		if err != nil {
			return fmt.Errorf("invalid file format: %w", err)
		}

		err = solveAndPrint(a, b, c)
		if err != nil {
			return err
		}
	}

	return nil
}

func readLine(line string) (a, b, c float64, err error) {
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
