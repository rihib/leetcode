//lint:file-ignore U1000 Ignore all unused code
package main

func generateParenthesis_iterative(n int) []string {
	var parentheses []string
	type state struct {
		parenthesis []byte
		open        int
		closed      int
	}
	stack := []state{{[]byte{}, 0, 0}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.open == n && current.closed == n {
			parentheses = append(parentheses, string(current.parenthesis))
			continue
		}
		if current.open < n {
			newParenthesis := append([]byte{}, current.parenthesis...)
			newParenthesis = append(newParenthesis, '(')
			stack = append(stack, state{newParenthesis, current.open + 1, current.closed})
		}
		if current.open > current.closed {
			newParenthesis := append([]byte{}, current.parenthesis...)
			newParenthesis = append(newParenthesis, ')')
			stack = append(stack, state{newParenthesis, current.open, current.closed + 1})
		}
	}
	return parentheses
}

func generateParenthesis_recursive(n int) []string {
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
