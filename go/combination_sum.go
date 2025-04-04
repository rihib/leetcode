//lint:file-ignore U1000 Ignore all unused code
package main

import "slices"

func combinationSumBacktrackingRecursion(candidates []int, target int) [][]int {
	var combinations [][]int
	var combination []int
	var generate func(int, int)
	generate = func(currentIndex, sum int) {
		if sum == target {
			newCombination := slices.Clone(combination)
			combinations = append(combinations, newCombination)
			return
		}
		if sum > target {
			return
		}
		for i := currentIndex; i < len(candidates); i++ {
			combination = append(combination, candidates[i])
			generate(i, sum+candidates[i])
			combination = combination[:len(combination)-1]
		}
	}
	generate(0, 0)
	return combinations
}

type frame struct {
	combination  []int
	currentIndex int
	sum          int
}

func combinationSumacktrackingIterative(candidates []int, target int) [][]int {
	var combinations [][]int
	stack := []frame{{[]int{}, 0, 0}}
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if f.sum == target {
			combinations = append(combinations, f.combination)
			continue
		}
		if f.sum > target {
			continue
		}
		for i := f.currentIndex; i < len(candidates); i++ {
			newCombination := slices.Clone(f.combination)
			newCombination = append(newCombination, candidates[i])
			stack = append(stack, frame{newCombination, i, f.sum + candidates[i]})
		}
	}
	return combinations
}

func combinationSumDP(candidates []int, target int) [][]int {
	combinationsGroup := make([][][]int, target+1)
	combinationsGroup[0] = [][]int{{}}
	for _, candidate := range candidates {
		for i := candidate; i <= target; i++ {
			for _, combination := range combinationsGroup[i-candidate] {
				newCombination := slices.Clone(combination)
				newCombination = append(newCombination, candidate)
				combinationsGroup[i] = append(combinationsGroup[i], newCombination)
			}
		}
	}
	return combinationsGroup[target]
}
