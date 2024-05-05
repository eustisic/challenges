package reader

import (
	"reflect"
	"testing"
)

func TestReadEncodedFileHeader(t *testing.T) {
	// Create a temporary test file for testing purposes

	// Define the test input and expected output
	fileName := "testfile.txt"
	expectedRuneData := map[rune]int{'a': 1, 'b': 2, 'c': 3}
	expectedPadding := 5

	// Call the function under test
	runeData, padding := ReadEncodedFileHeader(fileName)

	// Check if the returned rune data is as expected
	if !reflect.DeepEqual(runeData, expectedRuneData) {
		t.Errorf("Expected rune data %v, but got %v", expectedRuneData, runeData)
	}

	// Check if the returned padding value is as expected
	if padding != expectedPadding {
		t.Errorf("Expected padding value %d, but got %d", expectedPadding, padding)
	}

	// Add more test cases to cover other scenarios
}
