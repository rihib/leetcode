package blind75

func isValid(s string) bool {
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}
	var stack []rune

	for _, r := range s {
		v, found := brackets[r]

		if !found {
			stack = append(stack, r)
			continue
		}

		if len(stack) < 1 || stack[len(stack)-1] != v {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
