package main

import (
	"bufio"
	"fmt"
	"os"

	"kompressor/huffman"
	"kompressor/writer"
)

func main() {
	var reader *bufio.Reader
	var newFile []byte
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Please specify a file to pass to the kompressor")
	}

	filePath := args[0]

	file, err := os.Open(filePath)
	fileName := file.Name()
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader = bufio.NewReader(file)
	nodes, frequencies := huffman.MapCharacters(reader)

	node := huffman.BuildHuffmanTree(nodes)

	header := writer.Header(frequencies)
	// file := huffman.Encode(reader)

	writer.WriteFile(newFile, header, fileName)

	fmt.Println(node)
}

/*
[] sort the map by occurance
[] heapify the characters and create a priority queue

In this step your goal is to use the frequencies that you calculated in step 1 to build the binary tree. There’s a good explanation of how to do this complete with a visualisation here: https://opendsa-server.cs.vt.edu/ODSA/Books/CS3/html/Huffman.html

The examples used for the visualisation would be a good starting point for unit tests.
*/
