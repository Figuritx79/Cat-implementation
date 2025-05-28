package main

import (
	"bufio"
	"fmt"
	"os"
)

var input *bufio.Scanner

func main() {
	if len(os.Args) >= 2 {
		filePath := os.Args[1]
		file, err := os.Open(filePath)
		if err != nil {
			panic(1)
		}
		input = bufio.NewScanner(file)

		for input.Scan() {
			fmt.Println(input.Text())
		}
		file.Close()
	}

	input = bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
	}

}
