package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("Pass flags and file path")
	}

	if len(args) == 2 {
		filepath := args[1]

		lineCount := getLineCount(filepath)
		wordCount := getWordCount(filepath)
		byteCount := getBytesCount(filepath)
		fmt.Println(lineCount, wordCount, byteCount, filepath)
		return
	}
	flag := args[1]
	filepath := args[2]

	if flag == "-c" {
		bytes := getBytesCount(filepath)
		fmt.Println(bytes, filepath)
		return
	}

	if flag == "-l" {
		lines := getLineCount(filepath)
		fmt.Println(lines, filepath)
		return
	}

	if flag == "-w" {
		words := getWordCount(filepath)
		fmt.Println(words, filepath)

		return
	}

	if flag == "-m" {
		characters := getCharacterCount((filepath))
		fmt.Println(characters, filepath)
		return
	}

	return
}

func getBytesCount(filepath string) int {
	file, err := os.Stat(filepath)

	if err != nil {
		panic(err)
	}

	return int(file.Size())
}

func getLineCount(filepath string) int {
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

	return lineCount
}

func getWordCount(filepath string) int {
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

	return wordCount
}

func getCharacterCount(filepath string) int {
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

	return characterCount
}
