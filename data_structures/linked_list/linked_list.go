package linkedlist

import (
	"errors"
)

type LinkedList struct {
	firstNode *LinkedListNode
	lastNode  *LinkedListNode
}

type LinkedListNode struct {
	NextNode *LinkedListNode
	Value    interface{}
}

// Appends a node to the singly linked list
func (ll *LinkedList) Append(data interface{}) *LinkedList {
	if ll.firstNode == nil {
		ll.firstNode = &LinkedListNode{Value: data}
		ll.lastNode = ll.firstNode
		return ll
	}

	ll.lastNode.NextNode = &LinkedListNode{Value: data}
	ll.lastNode = ll.lastNode.NextNode
	return ll
}

// Prepends a node to the singly linked list
func (ll *LinkedList) Prepend(data interface{}) *LinkedList {
	newFirstNode := &LinkedListNode{Value: data, NextNode: ll.firstNode}
	ll.firstNode = newFirstNode
	return ll
}

// Find a node by 'index', index simply refers
// to how many nodes we must iterate through before
// finding the desired now
// If the node cannot be found, will return nil
func (ll *LinkedList) Get(index int) (*LinkedListNode, error) {
	if index < 0 {
		return nil, errors.New("index must be greater than or equal to 0")
	}

	currentNode := ll.firstNode
	count := 0

	for {
		if count == index {
			break
		}
		if currentNode != nil {
			currentNode = currentNode.NextNode
			count++
			continue
		}
		break
	}

	if currentNode != nil {
		return currentNode, nil
	}

	return nil, errors.New("passed an index out of range")
}

// Finds the data in the list and returns the index it was found at.
// If the node cannot be found, returns -1.
func (ll *LinkedList) Contains(data interface{}) int {
	count := 0
	currentNode := ll.firstNode
	found := false

	for {
		if currentNode.Value == data {
			found = true
			break
		}
		currentNode = currentNode.NextNode
		count++
		if currentNode == nil {
			break
		}
	}

	if found {
		return count
	}

	return -1

}

// Removes a node from the list at specified index
// Note, each previous index for an node will be -1.
// Returns a bool whether or not a node was removed
func (ll *LinkedList) Remove(index int) bool {
	count := 0
	currentNode := ll.firstNode

	if index < 0 {
		return false
	}

	if ll.firstNode == nil || ll.lastNode == nil {
		return false
	}

	if index > ll.Count()-1 {
		return false
	}

	for {
		if index == 0 {
			ll.firstNode = ll.firstNode.NextNode
			return true
		}
		if (index - 1) == count {
			newNextNode := currentNode.NextNode.NextNode
			currentNode.NextNode = newNextNode
			return true
		}
		currentNode = currentNode.NextNode
		count++
	}
}

// Counts the amount of linked nodes and returns
// the count.
func (ll *LinkedList) Count() int {
	count := 0
	currentNode := ll.firstNode
	for {
		if currentNode != nil {
			count++
			currentNode = currentNode.NextNode
		} else {
			break
		}
	}
	return count
}
