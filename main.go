package main

import (
	"bufio"
	"fmt"
	"os"
)

var input *bufio.Scanner

func main() {
	if len(os.Args) == 1 {
		runInteractiveMode()
		return
	}

	err := runFileMode(os.Args[1])
	if err != nil {
		fmt.Println("Error.", err)
		os.Exit(1)
	}
}
