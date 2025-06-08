package task1

func removeDuplicates(nums []int) int {
	var i int
	for j := 1; j < len(nums); j++ {
		if nums[i] < nums[j] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return i + 1
}
