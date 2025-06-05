package task1

func isValid(s string) bool {
	m := make(map[rune]int)

	for _, v := range s {
		switch v {
		case '(', '[', '{':
			m[v]++
		case ')':
			m['(']--
		case ']':
			m['[']--
		case '}':
			m['{']--
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
