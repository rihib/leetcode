//lint:file-ignore U1000 Ignore all unused code
package main

func generateParenthesis(n int) []string {
	var parentheses []string
	var stack []byte
	var generate func(int, int)
	generate = func(open int, closed int) {
		if open == n && closed == n {
			parentheses = append(parentheses, string(stack))
			return
		}
		if open < n {
			stack = append(stack, '(')
			generate(open+1, closed)
			stack = stack[:len(stack)-1]
		}
		if open > closed {
			stack = append(stack, ')')
			generate(open, closed+1)
			stack = stack[:len(stack)-1]
		}
	}
	generate(0, 0)
	return parentheses
}
