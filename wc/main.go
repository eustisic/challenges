package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filePath, flag string
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Please specify a file path or provide correct number of args")
		os.Exit(1)
	} else if len(args) == 1 {
		filePath = args[0]
	} else {
		flag = args[0]
		filePath = args[1]
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	lines, words, runes, bytes := getCounts(file)

	if len(flag) == 0 {
		fmt.Printf("%d %d %d %s\n", lines, words, bytes, file.Name())
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

func countBytes(f *os.File) int64 {
	info, err := f.Stat()
	if err != nil {
		fmt.Printf("Error getting file info: %s\n", err)
		os.Exit(1)
	}

	return info.Size()
}

func getCounts(f *os.File) (lines, words, runesCount, bytes int64) {
	bytes = countBytes(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		runesCount += int64(len(runes))
		lines++

		// Count words and runes in the line
		wordScanner := bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			words++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	return lines, words, runesCount, bytes
}
