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

type RuneStack struct {
	s []rune
}

func NewRuneStack(cap int) RuneStack {
	return RuneStack{
		s: make([]rune, 0, cap),
	}
}

func (s *RuneStack) Push(item rune) {
	s.s = append(s.s, item)
}

func (s *RuneStack) Pop() rune {
	var r rune
	if len(s.s) <= 0 {
		return r
	}

	r = s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return r
}

func (s *RuneStack) Length() int {
	return len(s.s)
}

func isValid2(s string) bool {
	rs := NewRuneStack(len(s))

	for _, v := range s {
		switch v {
		case '(', '{', '[':
			rs.Push(v)
		case ')':
			r := rs.Pop()
			if r != '(' {
				return false
			}
		case ']':
			r := rs.Pop()
			if r != '[' {
				return false
			}
		case '}':
			r := rs.Pop()
			if r != '{' {
				return false
			}
		}
	}

	if rs.Length() == 0 {
		return true
	} else {
		return false
	}
}
