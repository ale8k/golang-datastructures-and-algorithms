package main

import (
	"fmt"

	selectionsort "github.com/ale8k/golang_datastructures_and_algorithms/slice_sorting/selection_sort"
)

func main() {
	k := []int{5, 1, 7, 2}
	selectionsort.SelectionSortInt(k)
	fmt.Println(k)
}
