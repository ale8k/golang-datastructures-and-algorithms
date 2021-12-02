package selectionsort

// Sorts a slice of integers using selection sort algorithm
func SelectionSortInt(slice []int) {

	for i := 0; i < len(slice); i++ {
		minToSet := i
		// Use the iterations index for the inner loop
		// and increment by 1, so we can justcheck the rest of the array
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[minToSet] {
				// We know slice[j] < minimum to set, so we update our min to set index to j
				minToSet = j
			}
		}
		// Quickly check it has definitely changed
		if minToSet != i {
			temp := slice[minToSet]
			slice[minToSet] = slice[i]
			slice[i] = temp
		}

	}

}
