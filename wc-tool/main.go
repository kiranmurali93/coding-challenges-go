package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		return
	}
	flag := args[1]
	filepath := args[2]

	if flag == "-c" {
		file, err := os.Stat(filepath)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(file.Size(), filepath)
		return
	}

	if flag == "-l" {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		lineCount := 0
		for fileScanner.Scan() {
			lineCount++
		}

		fmt.Println(lineCount, filepath)
	}

	return
}
