//lint:file-ignore U1000 Ignore all unused code
package main

func generateParenthesisRecursive(n int) []string {
	var combinations []string
	currentCombination := make([]byte, 0, n*2)
	var generate func(int, int)
	generate = func(numLeft, numRight int) {
		if numLeft == n && numRight == n {
			combinations = append(combinations, string(currentCombination))
			return
		}
		if numLeft < n {
			currentCombination = append(currentCombination, '(')
			generate(numLeft+1, numRight)
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
		if numRight < numLeft {
			currentCombination = append(currentCombination, ')')
			generate(numLeft, numRight+1)
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}
	generate(0, 0)
	return combinations
}

func generateParenthesisIterative(n int) []string {
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
