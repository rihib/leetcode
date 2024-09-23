//lint:file-ignore U1000 Ignore all unused code
package main

func isValid(s string) bool {
	var stack []rune
	closeToOpens := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, bracket := range s {
		openBracket, ok := closeToOpens[bracket]
		if !ok {
			stack = append(stack, bracket)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != openBracket {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
