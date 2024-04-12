package writer

import (
	"reflect"
	"strings"
	"testing"
)

func TestGenerateHeaderFile(t *testing.T) {
	expectedHeader := "******begin******\n{\"97\":3,\"98\":2,\"99\":1}\n******end******"

	frequencies := map[rune]int{
		'a': 3,
		'b': 2,
		'c': 1,
	}

	actualHeader := string(Header(frequencies))

	if !reflect.DeepEqual(strings.TrimSpace(actualHeader), strings.TrimSpace(expectedHeader)) {
		t.Errorf("Generated header does not match expected header.\nExpected: %s\nActual: %s", expectedHeader, actualHeader)
	}
}
