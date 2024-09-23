//lint:file-ignore U1000 Ignore all unused code
package template

/*
Step1では再帰を使ってバックトラッキングを解いたので、スタックを使った方法も実装しました。
また、Step1の実装のリファクタもしました。
*/
func combinationSumBacktrackingStackStep4(candidates []int, target int) [][]int {
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

func combinationSumBacktrackingRecursionStep4(candidates []int, target int) [][]int {
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
