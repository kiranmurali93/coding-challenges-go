package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("Pass flags and file path")
	}
	flag := args[1]
	filepath := args[2]

	if flag == "-c" {
		file, err := os.Stat(filepath)

		if err != nil {
			panic(err)
		}

		fmt.Println(file.Size(), filepath)
		return
	}

	if flag == "-l" {
		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		lineCount := 0
		for fileScanner.Scan() {
			lineCount++
		}

		fmt.Println(lineCount, filepath)
		return
	}

	if flag == "-w" {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer file.Close()

		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanWords)

		wordCount := 0
		for fileScanner.Scan() {
			wordCount++
		}
		fmt.Println(wordCount, filepath)

		return
	}

	if flag == "-m" {
		file, err := os.Open(filepath)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanRunes)

		characterCount := 0
		for fileScanner.Scan() {
			characterCount++
		}
		fmt.Println(characterCount, filepath)
		return
	}

	return
}
