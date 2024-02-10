package blind75

func isValid(s string) bool {
	var stack []rune
	open_brackets := map[rune]struct{}{'(': {}, '{': {}, '[': {}}
	close_brackets := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, r := range s {
		if _, found := open_brackets[r]; found {
			stack = append(stack, r)
		}

		if v, found := close_brackets[r]; found {
			if len(stack) < 1 {
				return false
			}

			if stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
