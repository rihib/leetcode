//lint:file-ignore U1000 Ignore all unused code
package main

import "slices"

func generateParenthesisRecursive(n int) []string {
	var combinations []string
	var combination []rune
	var generate func(int, int)
	generate = func(open, close int) {
		if open == n && close == n {
			combinations = append(combinations, string(combination))
			return
		}
		if open < n {
			combination = append(combination, '(')
			generate(open+1, close)
			combination = combination[:len(combination)-1]
		}
		if open > close {
			combination = append(combination, ')')
			generate(open, close+1)
			combination = combination[:len(combination)-1]
		}
	}
	generate(0, 0)
	return combinations
}

type frame struct {
	combination []rune
	open        int
	close       int
}

func generateParenthesisIterative(n int) []string {
	var combinations []string
	stack := []frame{{[]rune{}, 0, 0}}
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if f.open == n && f.close == n {
			combinations = append(combinations, string(f.combination))
			continue
		}
		if f.open < n {
			newCombination := slices.Clone(f.combination)
			newCombination = append(newCombination, '(')
			stack = append(stack, frame{newCombination, f.open + 1, f.close})
		}
		if f.open > f.close {
			newCombination := slices.Clone(f.combination)
			newCombination = append(newCombination, ')')
			stack = append(stack, frame{newCombination, f.open, f.close + 1})
		}
	}
	return combinations
}
