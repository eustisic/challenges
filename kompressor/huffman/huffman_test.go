package huffman

import (
	"reflect"
	"testing"
)

func TestBuildHuffmanTree(t *testing.T) {
	nodes := []*Node{
		{char: 'a', freq: 3},
		{char: 'b', freq: 2},
		{char: 'c', freq: 1},
	}

	expectedRoot := &Node{
		char: '*',
		freq: 6,
		left: &Node{
			char: 'a',
			freq: 3,
		},
		right: &Node{
			char: '*',
			freq: 3,
			left: &Node{
				char: 'c',
				freq: 1,
			},
			right: &Node{
				char: 'b',
				freq: 2,
			},
		},
	}

	root := BuildHuffmanTree(nodes)

	if !reflect.DeepEqual(root, expectedRoot) {
		t.Errorf("Expected root node does not match actual root node")
	}
}

func TestDecode(t *testing.T) {
	nodes := []*Node{
		{char: 'a', freq: 3},
		{char: 'b', freq: 2},
		{char: 'c', freq: 1},
	}

	root := BuildHuffmanTree(nodes)

	decodedString := Decode("000111110", root)

	if decodedString != "aaabbc" {
		t.Errorf("Expected %s but got %s", "aaabbc", decodedString)
	}
}
