package main

import (
	"fmt"

	linkedlist "github.com/ale8k/golang_datastructures_and_algorithms/data_structures/linked_list"
)

func main() {
	list := &linkedlist.LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)

	val, err := list.Get(1)
	fmt.Println(val, err)
}
