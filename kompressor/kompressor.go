package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var reader *bufio.Reader
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Please specify a file to pass to the kompressor")
	}

	filePath := args[0]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}

	reader = bufio.NewReader(file)
	charMap := mapCharacters(reader)

	fmt.Printf("Character count a: %d\n", charMap['a'])
}

func mapCharacters(r *bufio.Reader) map[rune]int {
	charMap := make(map[rune]int)

	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		}

		for _, char := range line {
			charMap[char]++
		}
	}

	return charMap
}
