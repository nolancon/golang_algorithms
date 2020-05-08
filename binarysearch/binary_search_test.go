package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tcases := []struct{
		name string
		list []int
		number int
		found bool
	}{
		{
			name: "test 1",
			list: []int{1,2,3,4,5,6,7,8,9,10},
			number: 1,
			found: true,
		},
		{
			name: "test 2",
			list: []int{1,2,3,4,5,6,7,8,9,10},
			number: 10,
			found: true,
		},
		{
			name: "test 3",
			list: []int{34,56,134,158,456,577,633,776,881,910},
			number: 34,
			found: true,
		},
		{
			name: "test 4",
			list: []int{1,1,1,2,2,2,3,3,4,4,4,5,5,6,6,7},
			number: 4,
			found: true,
		},
		{
			name: "test 5",
			list: []int{1,1,1,1,1,1},
			number: 1,
			found: true,
		},
		{
			name: "test 6",
			list: []int{0,0,0,0},
			number: 1,
			found: false,
		},
		{
			name: "test 7",
			list: []int{0,0,0,0,0},
			number: 0,
			found: true,
		},

	}
	for _, tc := range tcases {
		found := BinarySearch(tc.list, tc.number)
		if found != tc.found {
			t.Errorf("Fail: %v - Expected found result to be %v, got %v", tc.name, tc.found, found)
		}
	}
}

