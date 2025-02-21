package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var input = bufio.NewScanner(os.Stdin)

func runInteractiveMode() {
	for {
		a := mustReadFloat("a = ")
		b := mustReadFloat("b = ")
		c := mustReadFloat("c = ")

		solveAndPrint(a, b, c)
		fmt.Println()
	}
}

func mustReadFloat(msg string) float64 {
	for {
		fmt.Print(msg)
		input.Scan()
		strF := input.Text()
		f, err := strconv.ParseFloat(strF, 64)
		if err != nil {
			fmt.Printf("Error. Expected a valid real number, got %s instead\n", strF)
			continue
		}
		return f
	}
}
