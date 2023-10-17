package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func printCountOf(countOf string, filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Failed to read the file.")
		return
	}

	scanner := bufio.NewScanner(file)
	if countOf == "bytes" {
		scanner.Split(bufio.ScanRunes)
	} else if countOf == "lines" {
		scanner.Split(bufio.ScanLines)
	} else if countOf == "words" {
		scanner.Split(bufio.ScanWords)
	} else if countOf == "characters" {
		scanner.Split(bufio.ScanBytes)
	} else {

	}

	count := 0
	for scanner.Scan() {
		count++
	}

	fmt.Printf("%s = %d\n", countOf, count)

}

func main() {

	var numberOfBytes, numberOfLines, numberOfWords, numberOfCharacters bool

	flag.BoolVar(&numberOfBytes, "c", false, "A boolean flag")
	flag.BoolVar(&numberOfLines, "l", false, "A boolean flag")
	flag.BoolVar(&numberOfWords, "w", false, "A boolean flag")
	flag.BoolVar(&numberOfCharacters, "m", false, "A boolean flag")

	flag.Parse()

	if !numberOfBytes && !numberOfLines && !numberOfWords && !numberOfCharacters {
		numberOfBytes = true
		numberOfLines = true
		numberOfWords = true
		numberOfCharacters = true
	}

	if flag.NArg() == 0 {
		fmt.Println("Please provide filename.")
		return
	}

	var secondParameter = flag.Arg(0)

	if numberOfLines {
		printCountOf("lines", secondParameter)
	}
	if numberOfWords {
		printCountOf("words", secondParameter)
	}
	if numberOfCharacters {
		printCountOf("bytes", secondParameter)
	}
	if numberOfBytes {
		printCountOf("characters", secondParameter)
	}
}
