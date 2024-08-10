//lint:file-ignore U1000 Ignore all unused code
package main

func combinationSum(candidates []int, target int) [][]int {
	var combinations [][]int
	var stack []int
	var findCombinations func(int, int)
	findCombinations = func(curr int, sum int) {
		if sum == target {
			combination := make([]int, len(stack))
			copy(combination, stack)
			combinations = append(combinations, combination)
			return
		}
		if sum > target {
			return
		}
		for i := curr; i < len(candidates); i++ {
			stack = append(stack, candidates[i])
			findCombinations(i, sum+candidates[i])
			stack = stack[:len(stack)-1]
		}
	}
	findCombinations(0, 0)
	return combinations
}
