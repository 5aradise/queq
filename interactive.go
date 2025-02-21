package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func runInteractiveMode() {
	input = bufio.NewScanner(os.Stdin)

	for {
		a := mustReadFloat("a = ")
		b := mustReadFloat("b = ")
		c := mustReadFloat("c = ")

		err := solveAndPrint(a, b, c)
		fmt.Println("Error.", err)
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
