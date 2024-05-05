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
const Prefix = "compressed_"

type Header struct {
	Frequencies string `json:"frequencies"`
	Padding     int    `json:"padding"`
}

func MarshalHeader(f map[rune]int, p int) []byte {
	frequencyString, err := json.Marshal(f)
	if err != nil {
		panic("Error marshalling frequencies")
	}

	h := Header{
		Frequencies: string(frequencyString),
		Padding:     p,
	}

	header, err := json.Marshal(h)
	if err != nil {
		panic("Error marshalling header")
	}

	return []byte(Begin + string(header) + "\n" + End)
}

func KompressToFile(file *os.File, f map[rune]int, fileName string, prefixCodes map[rune]string) {
	bitStrings := []string{}
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

	binaryString := strings.Join(bitStrings, "")

	// pad binary string to make a multiple of 8 then create byte slice to hold bytes
	// need to track length of pad and tack it to header so we know how to treat last byte
	padding := "00000000"[len(binaryString)%8:]
	paddedBinaryString := binaryString + padding
	packedBytes := make([]byte, len(paddedBinaryString)/8)

	// convert binary string to bytes
	for i := 0; i < len(paddedBinaryString); i += 8 {
		byteStr := paddedBinaryString[i : i+8]
		byteVal, _ := strconv.ParseUint(byteStr, 2, 8)
		packedBytes[i/8] = byte(byteVal)
	}

	newFileName := CreateFileName(Prefix + fileName)

	// Generate header with padding
	header := MarshalHeader(f, len(padding))

	err := os.WriteFile(newFileName, append(header, packedBytes...), 0644)

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
