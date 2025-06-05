package task1

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
