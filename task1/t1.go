package task1

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))

	for i, v := range nums {
		value, ok := numsMap[target-v]
		if ok {
			return []int{i, value}
		}
		numsMap[v] = i
	}

	return []int{}
}
