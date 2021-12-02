package bubblesort

// Sorts a slice of integers using sinking sort algorithm
func BubbleSortInt(slice []int) {
	swapWithNextIndex := func(i int, slice []int) {
		temp := slice[i+1]
		slice[i+1] = slice[i]
		slice[i] = temp
	}
	for i := range slice {
		for j := range slice {
			if i != len(slice)-1 && j != len(slice)-1 {
				// Iterate over every item for every item (n^2) and compare
				if slice[j] > slice[j+1] {
					swapWithNextIndex(j, slice)
				}
			}
		}
	}
}
