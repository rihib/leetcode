//lint:file-ignore U1000 Ignore all unused code
package blind75

func isValid(s string) bool {
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}
	var stack []rune

	for _, r := range s {
		v, ok := brackets[r]

		if !ok {
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
