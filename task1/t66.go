package task1

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		newDigit := plus(digits[i])
		digits[i] = newDigit
		if newDigit != 0 {
			return digits
		} else if newDigit+i == 0 {
			return append([]int{1}, digits...)
		} else {
			continue
		}
	}
	return []int{}
}

func plus(addend int) int {
	sum := addend + 1
	if sum < 10 {
		return sum
	} else {
		return 0
	}
}
