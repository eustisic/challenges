package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var filePath, flag, name string
	var reader *bufio.Reader
	args := os.Args[1:]

	if len(args) > 2 {
		fmt.Println("Please specify a file path and provide correct number of flags")
		os.Exit(1)

	} else if len(args) == 1 || len(args) == 2 {
		if len(args) == 1 {
			filePath = args[0]
		}
		if len(args) == 2 {
			flag = args[0]
			filePath = args[1]
		}

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()

		reader = bufio.NewReader(file)
		name = file.Name()

	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	lines, words, runes, bytes := getCounts(reader)

	if len(flag) == 0 {
		fmt.Printf("%d %d %d %s\n", lines, words, bytes, name)
		os.Exit(0)
	}

	switch args[0] {
	case "-c":
		fmt.Printf("%d %s\n", bytes, filePath)
	case "-l":
		fmt.Printf("%d %s\n", lines, filePath)
	case "-w":
		fmt.Printf("%d %s\n", words, filePath)
	case "-m":
		fmt.Printf("%d %s\n", runes, filePath)
	}
}

func getCounts(r *bufio.Reader) (lines, words, runes, bytes int) {

	for {
		line, err := r.ReadString('\n')
		bytes += len(line)
		runes += utf8.RuneCountInString(line)
		lines++

		// Count words and runes in the line
		wordScanner := bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			words++
		}

		if err == io.EOF {
			break
		}
	}

	return lines, words, runes, bytes
}
