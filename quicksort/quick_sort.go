package quicksort

func QuickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	// pivot is the number used to compare, use last number in slice 
	pivot := len(list)-1
	splitIndex := 0

	for i,_ := range list {
		if list[i] < list[pivot] {
			// swap number at split index with number at i 
			list[splitIndex], list[i] = list[i], list[splitIndex]
			splitIndex++
		}
	}

	// swap pivot number with number at split index
	list[pivot], list[splitIndex] = list[splitIndex], list[pivot]

	// now pivot is positioned correctly (numbers smaller than pivot on left, larger on right)
	QuickSort(list[: splitIndex])
	QuickSort(list[splitIndex+1 :])
	return list
}
