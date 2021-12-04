package heap

import (
	"fmt"
	"strconv"
)

// A basic int heap
// Accepts no size as we dynamically resize
// the heap in a slice fashion.
// Otherwise it is an ordinary MinHeap.
// Example:
// 		1
//	  5   3
//	9  6 8
// Arr form:
// 1 3 6 5 9 8
type MinHeap struct {
	nodes []int
}

// Gets parent index for a given node
func (h *MinHeap) ParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

// Get left child index for a given node index
func (h *MinHeap) LeftChildIndex(parentNodeIndex int) int {
	return parentNodeIndex*2 + 1
}

// Get right child index for a given node index
func (h *MinHeap) RightChildIndex(parentNodeIndex int) int {
	return h.LeftChildIndex(parentNodeIndex) + 1
}

// Insert a node into the heap
func (h *MinHeap) InsertNode(node ...int) *MinHeap {
	for _, v := range node {
		h.nodes = append(h.nodes, v)
		nodeIndex := len(h.nodes) - 1

		for nodeIndex != 0 && h.nodes[nodeIndex] < h.nodes[h.ParentIndex(nodeIndex)] {
			fmt.Println("Shuffling node value: ", node)
			h.swapNodes(nodeIndex, h.ParentIndex(nodeIndex))
			nodeIndex = h.ParentIndex(nodeIndex)
		}
	}
	return h
}

func (h *MinHeap) String() string {
	initial := "[["
	for _, v := range h.nodes {
		initial += strconv.Itoa(v)
		initial += ", "
	}
	initial += " ]]"
	return initial
}

// Swaps two nodes around
func (h *MinHeap) swapNodes(i1 int, i2 int) {
	temp := h.nodes[i1]
	h.nodes[i1] = h.nodes[i2]
	h.nodes[i2] = temp
}
