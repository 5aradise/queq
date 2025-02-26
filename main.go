package main

import (
	"fmt"
	"os"
)

func main() {
	var mode mode

	if len(os.Args) == 1 {
		mode = newInteractiveMode(os.Stdin, os.Stdout)
		err := mode.run()
		if err != nil {
			fmt.Println("Error.", err)
			os.Exit(1)
		}
	} else {
		filePath := os.Args[1]
		f, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error.", err)
			os.Exit(1)
		}
		defer f.Close()

		mode = newFileMode(f, os.Stdout)
		err = mode.run()
		if err != nil {
			fmt.Println("Error.", err)
			os.Exit(1)
		}
	}
}
