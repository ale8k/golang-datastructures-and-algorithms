package interpolationsearch

// Runs a search based on interpolation
func InterpolationSearch(param int, slice []int) int {
	if param < 0 {
		return -1
	}
	low := 0
	mid := -1
	high := len(slice) - 1

	for {

		if low == high || slice[low] == slice[high] {
			return -1
		}

		//	 		((Hi -  Lo)/(slice[Hi]  -slice[Lo]     ))*(X-slice[Lo])
		mid = low + ((high-low)/(slice[high]-slice[low]))*(param-slice[low])

		// where −
		// 	A    = list
		// 	Lo   = Lowest index of the list
		// 	Hi   = Highest index of the list
		// 	A[n] = Value stored at index n in the list

		if slice[mid] == param {
			return mid
		} else {
			if slice[mid] < param {
				low = mid + 1
			} else if slice[mid] > param {
				high = mid - 1
			}
		}
	}
}
