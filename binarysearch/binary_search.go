package binarysearch

func BinarySearch(list []int, num int) bool {
	lo := 0
	hi := len(list)-1

	for lo <= hi {
		mid := (lo + hi)/2

		if num == list[mid] {
			return true
		}else if num > list[mid] {
			lo = mid + 1
		}else  if num < list[mid] {
			hi = mid - 1
		}
	}

	return false
}


