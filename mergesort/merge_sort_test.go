package mergesort

import (
	"testing"
	"reflect"
)

func TestMergeSort(t *testing.T) {
	tcases := []struct{
		name string
		unsorted []int
		sorted []int
	}{
		{
			name: "test 1",
			unsorted: []int{10,9,8,7,6,5,4,3,2,1},
			sorted: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},

		},
		{
			name: "test 2",
			unsorted: []int{34, 26, 134, 58, 45, 77, 33, 17, 1, 0},
			sorted: []int{0, 1, 17, 26, 33, 34, 45, 58, 77, 134},
		},
		{
			name: "test 3",
			unsorted: []int{1, 1, 1, 0, 0, 0, 0},
			sorted: []int{0, 0, 0, 0, 1, 1, 1},
		},

	}
	for _, tc := range tcases {
		sorted := MergeSort(tc.unsorted)
		if !reflect.DeepEqual(sorted, tc.sorted) {
			t.Errorf("Fail: %v - Expected sorted to be %v, got %v", tc.name, tc.sorted, sorted)
		}
	}
}

