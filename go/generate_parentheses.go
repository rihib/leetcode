//lint:file-ignore U1000 Ignore all unused code
package main

import "slices"

func generateParenthesisRecursive(n int) []string {
	var combinations []string
	combination := make([]rune, 0, n*2)
	var generate func(int, int)
	generate = func(openNum, closeNum int) {
		if openNum == n && closeNum == n {
			combinations = append(combinations, string(combination))
			return
		}
		if openNum < n {
			combination = append(combination, '(')
			generate(openNum+1, closeNum)
			combination = combination[:len(combination)-1]
		}
		if closeNum < openNum {
			combination = append(combination, ')')
			generate(openNum, closeNum+1)
			combination = combination[:len(combination)-1]
		}
	}
	generate(0, 0)
	return combinations
}

type frame struct {
	combination []rune
	openNum     int
	closeNum    int
}

func generateParenthesisIterative(n int) []string {
	var combinations []string
	stack := []frame{{make([]rune, 0, n*2), 0, 0}}
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if f.openNum == n && f.closeNum == n {
			combinations = append(combinations, string(f.combination))
			continue
		}
		if f.openNum < n {
			newCombination := append(slices.Clone(f.combination), '(')
			stack = append(stack, frame{newCombination, f.openNum + 1, f.closeNum})
		}
		if f.closeNum < f.openNum {
			newCombination := append(slices.Clone(f.combination), ')')
			stack = append(stack, frame{newCombination, f.openNum, f.closeNum + 1})
		}
	}
	return combinations
}
