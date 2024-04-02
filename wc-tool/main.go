package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type FlagOptions struct {
	bytes      bool
	lines      bool
	words      bool
	characters bool
}

type FileOutput struct {
	bytes      int
	lines      int
	words      int
	characters int
}

func main() {

	var flags FlagOptions

	flag.BoolVar(&flags.bytes, "c", false, "count bytes")
	flag.BoolVar(&flags.lines, "l", false, "count lines")
	flag.BoolVar(&flags.words, "w", false, "count words")
	flag.BoolVar(&flags.characters, "m", false, "count runes")
	flag.Parse()

	fileNames := flag.CommandLine.Args()

	if !flags.bytes && !flags.lines && !flags.words && !flags.characters {
		flags.bytes = true
		flags.words = true
		flags.lines = true
	}

	parseFilesAndPrintData(fileNames, flags)

}

func parseFilesAndPrintData(fileName []string, flags FlagOptions) {
	var total FileOutput
	for _, filepath := range fileName {
		var output FileOutput

		if flags.bytes {
			bytes := getBytesCount(filepath)
			output.bytes = bytes
			total.bytes = total.bytes + bytes
		}

		if flags.lines {
			lines := getLineCount(filepath)
			output.lines = lines
			total.lines = total.lines + lines
		}

		if flags.words {
			words := getWordCount(filepath)
			output.words = words
			total.words = total.words + words
		}

		if flags.characters {
			characters := getCharacterCount((filepath))
			output.characters = characters
			total.characters = total.characters + characters
		}

		fmt.Println(parseOutput(output, flags), filepath)
	}
	if len(fileName) > 1 {
		fmt.Println(parseOutput(total, flags), "total")
	}

}

func parseOutput(outputData FileOutput, flags FlagOptions) string {
	var output string
	if flags.lines {
		output = output + " " + strconv.Itoa(outputData.lines)
	}

	if flags.words {
		output = output + " " + strconv.Itoa(outputData.words)
	}

	if flags.characters {
		output = output + " " + strconv.Itoa(outputData.characters)
	}

	if flags.bytes {
		output = output + " " + strconv.Itoa(outputData.bytes)
	}

	return output

}

func getBytesCount(filepath string) int {
	file, err := os.Stat(filepath)

	if err != nil {
		log.Fatal(err)
	}

	return int(file.Size())
}

func getLineCount(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
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
		log.Fatal(err)
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
