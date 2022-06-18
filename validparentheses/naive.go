package validparentheses

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, r := range s {
		if len(stack) == 0 {
			stack = append(stack, r)
			continue
		}
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
			if r-1 == stack[len(stack)-1] || r-2 == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
