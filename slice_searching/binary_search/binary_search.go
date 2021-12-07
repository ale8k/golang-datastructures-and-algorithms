package binarysearch

// Runs a binary search on a given slice numerical values
func BinarySearch(param int, slice []int) int {
	low := 0
	high := len(slice) - 1
	mid := low + (high-low)/2

	for {
		if mid > len(slice)-1 || mid < 0 {
			return -1
		} else if slice[mid] > param {
			mid -= 1
		} else if slice[mid] < param {
			mid += 1
		} else {
			return mid
		}
	}

}
