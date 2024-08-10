//lint:file-ignore U1000 Ignore all unused code
package main

func combinationSum(candidates []int, target int) [][]int {
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
