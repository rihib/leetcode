//lint:file-ignore U1000 Ignore all unused code
package main

func generateParenthesis(n int) []string {
	var parentheses []string
	parenthesis := make([]byte, 0, n*2)
	var generate func(int, int)
	generate = func(open int, closed int) {
		if open == n && closed == n {
			parentheses = append(parentheses, string(parenthesis))
			return
		}
		if open < n {
			parenthesis = append(parenthesis, '(')
			generate(open+1, closed)
			parenthesis = parenthesis[:len(parenthesis)-1]
		}
		if open > closed {
			parenthesis = append(parenthesis, ')')
			generate(open, closed+1)
			parenthesis = parenthesis[:len(parenthesis)-1]
		}
	}
	generate(0, 0)
	return parentheses
}
