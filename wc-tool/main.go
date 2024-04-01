package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Options struct {
	bytes      bool
	lines      bool
	words      bool
	characters bool
}

func main() {

	var flags Options

	flag.BoolVar(&flags.bytes, "c", false, "count bytes")
	flag.BoolVar(&flags.lines, "l", false, "count lines")
	flag.BoolVar(&flags.words, "w", false, "count words")
	flag.BoolVar(&flags.characters, "m", false, "count runes")
	flag.Parse()

	fileNames := flag.CommandLine.Args()
	// TODO: Add support for multiple files
	filepath := fileNames[0]

	if !flags.bytes || !flags.lines || flags.words || flags.characters {

		lineCount := getLineCount(filepath)
		wordCount := getWordCount(filepath)
		byteCount := getBytesCount(filepath)
		fmt.Println(lineCount, wordCount, byteCount, filepath)
		return
	}

	// TODO: add formatting for multiple flags
	if flags.bytes {
		bytes := getBytesCount(filepath)
		fmt.Println(bytes, filepath)
	}

	if flags.lines {
		lines := getLineCount(filepath)
		fmt.Println(lines, filepath)
	}

	if flags.words {
		words := getWordCount(filepath)
		fmt.Println(words, filepath)

	}

	if flags.characters {
		characters := getCharacterCount((filepath))
		fmt.Println(characters, filepath)
	}

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
