//lint:file-ignore U1000 Ignore all unused code
package main

func combinationSumDP(candidates []int, target int) [][]int {
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

func combinationSumBacktrackingIterative(candidates []int, target int) [][]int {
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
			combinations = append(combinations, current.combination)
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

func combinationSumBacktrackingRecursion(candidates []int, target int) [][]int {
	var combinations [][]int
	var combination []int
	var generateCombinations func(int, int)
	generateCombinations = func(currentIndex int, sum int) {
		if sum == target {
			combinations = append(combinations, append([]int{}, combination...))
			return
		}
		if sum > target {
			return
		}
		for i := currentIndex; i < len(candidates); i++ {
			combination = append(combination, candidates[i])
			generateCombinations(i, sum+candidates[i])
			combination = combination[:len(combination)-1]
		}
	}
	generateCombinations(0, 0)
	return combinations
}
