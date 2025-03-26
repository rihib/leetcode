//lint:file-ignore U1000 Ignore all unused code
package main

import "slices"

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

type combination struct {
	parenthesis []byte
	numLeft     int
	numRight    int
}

func generateParenthesisIterative(n int) []string {
	var combinations []string
	stack := []combination{{[]byte{}, 0, 0}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.numLeft == n && current.numRight == n {
			combinations = append(combinations, string(current.parenthesis))
			continue
		}
		if current.numLeft < n {
			newParenthesis := slices.Clone(current.parenthesis)
			newParenthesis = append(newParenthesis, '(')
			stack = append(stack, combination{newParenthesis, current.numLeft + 1, current.numRight})
		}
		if current.numRight < current.numLeft {
			newParenthesis := slices.Clone(current.parenthesis)
			newParenthesis = append(newParenthesis, ')')
			stack = append(stack, combination{newParenthesis, current.numLeft, current.numRight + 1})
		}
	}
	return combinations
}
