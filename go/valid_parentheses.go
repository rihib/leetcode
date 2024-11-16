//lint:file-ignore U1000 Ignore all unused code
package main

func isValid(s string) bool {
	closeToOpen := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []rune
	for _, r := range s {
		p, ok := closeToOpen[r]
		if !ok {
			stack = append(stack, r)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != p {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
