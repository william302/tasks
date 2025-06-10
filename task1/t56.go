package task1

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	func() {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
	}()

	var i int = 0
	for j := 1; j < len(intervals); j++ {
		left, right := intervals[i], intervals[j]
		meg := func(left, right []int) ([]int, bool) {
			merged := make([]int, 2)
			if left[1] < right[0] {
				return []int{}, false
			} else {
				merged[0] = min(left[0], right[0])
				merged[1] = max(left[1], right[1])
			}
			return merged, true
		}
		merged, ok := meg(left, right)
		if ok {
			intervals[i] = merged
		} else {
			i++
			intervals[i] = right
		}
	}
	return intervals[:i+1]
}
