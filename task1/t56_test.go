package task1

import "testing"

func TestT56(t *testing.T) {
	intervals := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}

	merge(intervals)
}
