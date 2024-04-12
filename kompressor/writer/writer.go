package writer

import (
	"encoding/json"
	"os"
)

func Header(frequencies map[rune]int) []byte {
	begin := "******begin******\n"
	end := "******end******\n"

	header, err := json.Marshal(frequencies)
	if err != nil {
		panic("Error marshalling frequencies")
	}

	return []byte(begin + string(header) + "\n" + end)
}

func WriteFile(file []byte, header []byte, fileName string) {
	err := os.WriteFile("compressed_"+fileName, append(header, file...), 0644)

	if err != nil {
		panic(err.Error())
	}
}

// func ReadFile() {

// }
