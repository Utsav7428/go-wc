package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings" 
)

func main() {
	// Define command-line flags for line, word, and byte count
	countLines := flag.Bool("l", false, "Count lines")
	countWords := flag.Bool("w", false, "Count words")
	countBytes := flag.Bool("c", false, "Count bytes")
	flag.Parse()

	// If no flags are provided, count everything by default
	if !*countLines && !*countWords && !*countBytes {
		*countLines = true
		*countWords = true
		*countBytes = true
	}

	// Determine input source: file or stdin
	var reader io.Reader
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	// Create scanner and counters
	scanner := bufio.NewScanner(reader)
	lineCount := 0
	wordCount := 0
	byteCount := 0

	// Set up scanner for line and word counting
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		byteCount += len(line) + 1 // +1 for newline

		if *countWords {
			wordScanner := bufio.NewScanner(strings.NewReader(line))
			wordScanner.Split(bufio.ScanWords)
			for wordScanner.Scan() {
				wordCount++
			}
		}
	}

	// Print results as per flags
	if *countLines {
		fmt.Printf("Lines: %d\n", lineCount)
	}
	if *countWords {
		fmt.Printf("Words: %d\n", wordCount)
	}
	if *countBytes {
		fmt.Printf("Bytes: %d\n", byteCount)
	}
}

/*
Example usage:
1. go run wc.go -l test.txt          // only count lines
2. go run wc.go -w -c test.txt      // count words and bytes
3. cat test.txt | go run wc.go      // read from stdin and count all
*/
