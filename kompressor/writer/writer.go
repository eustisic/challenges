package writer

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
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

func WriteFile(file *os.File, header []byte, fileName string, prefixCodes map[rune]string) {
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

	err := os.WriteFile("compressed_"+fileName, append(header, newFile...), 0644)

	if err != nil {
		panic(err.Error())
	}
}

// func ReadFile() {

// }

/*
This is the moment you’ve been building up to! In this step your goal is to encode the text using the code table and write the compressed content of the source text to the output file file, appending it after the header. Don’t forget translate the prefixes into bit strings and pack them into bytes to achieve the compression.
*/
