package writer

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestGenerateHeaderFile(t *testing.T) {
	expectedHeader := "******begin******\n{\"frequencies\":\"{\\\"97\\\":3,\\\"98\\\":2,\\\"99\\\":1}\",\"padding\":7}\n******end******"

	frequencies := map[rune]int{
		'a': 3,
		'b': 2,
		'c': 1,
	}

	actualHeader := string(MarshalHeader(frequencies, 7))

	if !reflect.DeepEqual(strings.TrimSpace(actualHeader), strings.TrimSpace(expectedHeader)) {
		t.Errorf("Generated header does not match expected header.\nExpected: %s\nActual: %s", expectedHeader, actualHeader)
	}
}

func TestCreateFileName(t *testing.T) {
	// Test case where file does not exist
	filename := "testfile.txt"
	expectedNewFileName := "testfile_1.txt"
	_, err := os.Create(filename)
	if err != nil {
		t.Errorf("Error in test setup")
	}
	newFilename := CreateFileName(filename)
	defer os.Remove(filename) // Cleanup after the test

	if newFilename != expectedNewFileName {
		t.Errorf("Expected newfile to be named %s, but got %s", expectedNewFileName, newFilename)
	}
	defer os.Remove(newFilename) // Cleanup after the test
}
