//lint:file-ignore U1000 Ignore all unused code
package main

func combinationSum_dp(candidates []int, target int) [][]int {
	combinationsGroups := make([][][]int, target+1)
	combinationsGroups[0] = [][]int{{}}
	for _, candidate := range candidates {
		for i := candidate; i <= target; i++ {
			for _, combination := range combinationsGroups[i-candidate] {
				newCombination := append([]int{}, combination...)
				newCombination = append(newCombination, candidate)
				combinationsGroups[i] = append(combinationsGroups[i], newCombination)
			}
		}
	}
	return combinationsGroups[target]
}

func combinationSum_backtracking_stack(candidates []int, target int) [][]int {
	combinations := [][]int{}
	type state struct {
		combination []int
		sum         int
		index       int
	}
	stack := []state{{[]int{}, 0, 0}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.sum == target {
			combinations = append(combinations, append([]int{}, current.combination...))
			continue
		}
		for i := current.index; i < len(candidates); i++ {
			newSum := current.sum + candidates[i]
			if newSum > target {
				continue
			}
			newCombination := append([]int{}, current.combination...)
			newCombination = append(newCombination, candidates[i])
			stack = append(stack, state{newCombination, newSum, i})
		}
	}
	return combinations
}

func combinationSum_backtracking_recursion(candidates []int, target int) [][]int {
	var combinations [][]int
	var stack []int
	var generateCombinations func(int, int)
	generateCombinations = func(currentIndex int, sum int) {
		if sum == target {
			combinations = append(combinations, append([]int{}, stack...))
			return
		}
		if sum > target {
			return
		}
		for i := currentIndex; i < len(candidates); i++ {
			stack = append(stack, candidates[i])
			generateCombinations(i, sum+candidates[i])
			stack = stack[:len(stack)-1]
		}
	}
	generateCombinations(0, 0)
	return combinations
}
