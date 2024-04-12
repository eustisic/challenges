package huffman

import (
	"bufio"
	"container/heap"
	"io"
)

// Node represents a node in the Huffman Tree.
type Node struct {
	char  rune
	freq  int
	left  *Node
	right *Node
}

// Item represents an item in the priority queue.
type Item struct {
	value    *Node
	priority int
	index    int
}

// PriorityQueue implements a priority queue based on a min-heap.
type PriorityQueue []*Item

// Implement heap.Interface methods for PriorityQueue

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func MapCharacters(r *bufio.Reader) ([]*Node, map[rune]int) {
	charMap := make(map[rune]int)
	nodes := []*Node{}

	for {
		char, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		charMap[char]++
	}

	for char, freq := range charMap {
		node := &Node{char: char, freq: freq}
		nodes = append(nodes, node)
	}

	return nodes, charMap
}

func BuildHuffmanTree(nodes []*Node) *Node {
	pq := make(PriorityQueue, len(nodes))

	for i, node := range nodes {
		pq[i] = &Item{
			value:    node,
			priority: node.freq,
			index:    i,
		}
	}

	heap.Init(&pq)

	for len(pq) > 1 {
		min1 := heap.Pop(&pq).(*Item).value
		min2 := heap.Pop(&pq).(*Item).value

		parent := &Node{char: '*', freq: min1.freq + min2.freq, left: min1, right: min2}
		heap.Push(&pq, &Item{value: parent, priority: parent.freq})
	}

	return heap.Pop(&pq).(*Item).value
}

func Decode(input string, root *Node) string {
	currentNode := root
	output := ""

	for _, bit := range input {
		if bit == '0' {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}

		if currentNode.char != '*' {
			output += string(currentNode.char)
			currentNode = root
		}
	}
	return output
}

func BuildPrefixCodeTable(root *Node, code string, prefixCodes map[string]rune) {
	if root == nil {
		return
	}

	if root.char != '*' {
		prefixCodes[code] = root.char
	}

	BuildPrefixCodeTable(root.left, code+"0", prefixCodes)
	BuildPrefixCodeTable(root.right, code+"1", prefixCodes)
}
