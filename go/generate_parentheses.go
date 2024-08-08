//lint:file-ignore U1000 Ignore all unused code
package main

func generateParenthesis(n int) []string {
	var parentheses []string
	var stack []byte
	var backtrack func(int, int)
	backtrack = func(open int, closed int) {
		if open == n && closed == n {
			parentheses = append(parentheses, string(stack))
			return
		}
		if open < n {
			stack = append(stack, '(')
			backtrack(open+1, closed)
			stack = stack[:len(stack)-1]
		}
		if closed < open {
			stack = append(stack, ')')
			backtrack(open, closed+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)
	return parentheses
}
