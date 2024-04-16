package main

import (
	"bufio"
	"os"

	"kompressor/huffman"
	"kompressor/reader"
	"kompressor/writer"
)

func main() {
	var r *bufio.Reader
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

	// read file and map to characters
	r = bufio.NewReader(file)
	nodes, frequencies := huffman.MapCharacters(r)

	// build the huffman tree from frequencies
	node := huffman.BuildHuffmanTree(nodes)

	// generate the header for decoding
	header := writer.Header(frequencies)

	// generate prefix key
	prefixCodes := map[rune]string{}
	huffman.BuildPrefixCodeTable(node, "", prefixCodes)

	// reset pointer to begining of file
	file.Seek(0, 0)
	// write to file given the prefix key
	writer.WriteFile(file, header, fileName, prefixCodes)

	reader.ReadFile("compressed_" + fileName)
}

/*
[] sort the map by occurance
[] heapify the characters and create a priority queue

In this step your goal is to use the frequencies that you calculated in step 1 to build the binary tree. There’s a good explanation of how to do this complete with a visualisation here: https://opendsa-server.cs.vt.edu/ODSA/Books/CS3/html/Huffman.html

The examples used for the visualisation would be a good starting point for unit tests.
*/
