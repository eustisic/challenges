package writer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const Begin = "******begin******\n"
const End = "******end******\n"

func Header(frequencies map[rune]int) []byte {
	header, err := json.Marshal(frequencies)
	if err != nil {
		panic("Error marshalling frequencies")
	}

	return []byte(Begin + string(header) + "\n" + End)
}

func KompressToFile(file *os.File, header []byte, fileName string, prefixCodes map[rune]string) {
	bitStrings := []string{}
	newFile := []byte{}
	reader := bufio.NewReader(file)

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		bitStrings = append(bitStrings, prefixCodes[char])
	}

	concatenatedBinary := strings.Join(bitStrings, "")

	var packedByte string
	for _, char := range concatenatedBinary {
		if len(packedByte) < 8 {
			packedByte += string(char)
			continue
		}

		newByte, _ := strconv.ParseUint(packedByte, 2, 8)
		newFile = append(newFile, byte(newByte))
		packedByte = ""
	}

	if len(packedByte) > 0 {
		newByte, _ := strconv.ParseUint(packedByte, 2, 8)
		newFile = append(newFile, byte(newByte))
	}

	newFileName := CreateFileName("compressed_" + fileName)

	err := os.WriteFile(newFileName, append(header, newFile...), 0644)

	if err != nil {
		panic(err.Error())
	}
}

func CreateFileName(filename string) string {
	if _, err := os.Stat(filename); err == nil {
		baseName, extension := splitFileName(filename)
		index := 1
		for {
			newFilename := fmt.Sprintf("%s_%d%s", baseName, index, extension)
			if _, err := os.Stat(newFilename); os.IsNotExist(err) {
				filename = newFilename
				break
			}
			index++
		}
	}

	return filename
	// Write content to the file if needed
}

func splitFileName(filename string) (string, string) {
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	return name, ext
}
