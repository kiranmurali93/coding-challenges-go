package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
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

	if len(fileNames) < 1 {
		reader := bufio.NewReader(os.Stdin)
		parseFileFromIo(reader, flags)
		return
	}

	parseFileIfFilePathExists(fileNames, flags)

}

func parseFileIfFilePathExists(fileName []string, flags FlagOptions) {
	var total FileOutput
	for _, filepath := range fileName {
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		reader := bufio.NewReader(file)

		fileStats := getFileStats(reader)

		fmt.Println(parseOutput(fileStats, flags), filepath)
		total.bytes += fileStats.bytes
		total.characters += fileStats.characters
		total.lines += fileStats.lines
		total.words += fileStats.words
	}
	if len(fileName) > 1 {
		fmt.Println(parseOutput(total, flags), "total")
	}

}

func parseFileFromIo(reader *bufio.Reader, flags FlagOptions) {
	fileStats := getFileStats(reader)

	statsOutput := parseOutput(fileStats, flags)

	fmt.Println(statsOutput)
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

func getFileStats(reader *bufio.Reader) FileOutput {
	var fileStats FileOutput
	var prevCharacter rune

	for {
		character, bytes, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				// to handle case where eof appears as soon as the word without space
				if prevCharacter != rune(0) && !unicode.IsSpace(prevCharacter) {
					fileStats.words++
				}
				break
			}
			log.Fatal(err)
		}

		fileStats.bytes += int(bytes)
		fileStats.characters++

		if !unicode.IsSpace(prevCharacter) && unicode.IsSpace(character) {
			fileStats.words++
		}

		if character == '\n' {
			fileStats.lines++
		}

		prevCharacter = character
	}

	return fileStats
}
