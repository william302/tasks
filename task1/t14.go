package task1

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	firstStr := strs[0]
	strs = strs[1:]

	var count int
	for i, v := range firstStr {
		match := func() bool {
			for _, str := range strs {
				if []rune(str)[i] != v {
					return false
				}
			}
			return true
		}()

		if !match {
			break
		} else {
			count++
		}
	}

	return string([]rune(firstStr)[:count])
}
