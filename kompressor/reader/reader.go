package reader

import (
	"bytes"
	"encoding/json"
	"kompressor/writer"
	"os"
	"strconv"
)

// reads uncompressed file
func ReadFile(fileName string) map[rune]int {
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

	data := make(map[string]int)
	err = json.Unmarshal(encodingSection, &data)
	if err != nil {
		panic("Error parsing JSON")
	}

	// Convert the string keys to rune keys
	runeData := make(map[rune]int)
	for k, v := range data {
		num, _ := strconv.ParseInt(k, 10, 8)
		runeData[rune(num)] = v
	}

	return runeData
}
