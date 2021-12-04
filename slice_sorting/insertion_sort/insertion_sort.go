package insertionsort

// Sorts a slice of integers using insertion sort algorithm
func InsertionSortInt(slice []int) {
	for i := 0; i < len(slice); i++ {
		currentIndex := i + 1
		for currentIndex != len(slice) && currentIndex != 0 && slice[currentIndex] < slice[currentIndex-1] {
			tempMinus := slice[currentIndex-1]
			slice[currentIndex-1] = slice[currentIndex]
			slice[currentIndex] = tempMinus
			currentIndex--
		}
	}
}
