package quicksort

func QuickSort(list []int) []int {
	start := 0
	end := len(list)-1
	return sort(list, start, end)
}

func sort(list []int, start int, end int) []int {
	if end - start < 1 {
		return list
	}

	// pivot is the number used to compare, use last number in slice 
	pivot := list[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if list[i] < pivot {

			// swap number at split index with number at i 
			tmp := list[splitIndex]
			list[splitIndex] = list[i]
			list[i] = tmp

			splitIndex++
		}
	}

	// swap pivot number with number at split index
	list[end] = list[splitIndex]
	list[splitIndex] = pivot
	// now pivot is positioned correctly (numbers smaller than pivot on left, larger on right) eg 4 -> {2,2,3,1,4,6,5,7,8}

	sort(list, start, splitIndex-1)
	sort(list, splitIndex+1, end)
	return list
}
