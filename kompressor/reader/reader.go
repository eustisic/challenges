package reader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kompressor/huffman"
	"kompressor/writer"
	"os"
	"strconv"
	"strings"
)

// reads uncompressed file
func ReadEncodedFileHeader(fileName string) (map[rune]int, int) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}

	begin := writer.Begin
	end := writer.End

	beginIdx := bytes.Index(file, []byte(begin))
	endIdx := bytes.Index(file, []byte(end))
	if beginIdx == -1 || endIdx == -1 {
		panic("Section not found.")
	}

	encodingSection := file[beginIdx+len(begin) : endIdx]

	var data writer.Header
	err = json.Unmarshal(encodingSection, &data)
	if err != nil {
		panic("Error parsing JSON")
	}

	var frequencies map[string]int
	err = json.Unmarshal([]byte(data.Frequencies), &frequencies)
	if err != nil {
		panic("Error parsing frequencies JSON")
	}

	// Convert the string keys to rune keys
	runeData := make(map[rune]int)
	for k, v := range frequencies {
		num, _ := strconv.ParseInt(k, 10, 8)
		runeData[rune(num)] = v
	}

	return runeData, data.Padding
}

func ReadAndDecode(fileName string, root *huffman.Node, p int) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}

	end := writer.End

	endIdx := bytes.Index(file, []byte(end)) + len(end)

	byteString := ""
	for i := endIdx; i < len(file); i++ {
		str := byteToBinaryString(file[i])
		if i == len(file)-1 {
			str = str[:(8 - p)]
		}
		byteString += str
	}

	output := huffman.DecodeString(byteString, root)

	newFileName := writer.CreateFileName(strings.TrimPrefix(fileName, writer.Prefix))

	os.WriteFile(newFileName, []byte(output), 0644)
}

func byteToBinaryString(b byte) string {
	return fmt.Sprintf("%08b", b)
}
