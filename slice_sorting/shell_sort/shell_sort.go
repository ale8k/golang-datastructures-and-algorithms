package shellsort

func ShellSort(slice []int) {
	interval := 1
	inner := 0
	outer := 0
	insertValue := 0

	for interval <= len(slice)/3 {
		interval = interval*3 + 1
	}

	for interval > 0 {
		for outer = interval; outer < len(slice); outer++ {
			insertValue = slice[outer]
			inner = outer

			for inner > interval-1 && slice[inner-interval] >= insertValue {
				slice[inner] = slice[inner-interval]
				inner = inner - interval
			}

			slice[inner] = insertValue
		}
		interval = (interval - 1) / 3
	}

}
