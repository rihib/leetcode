//lint:file-ignore U1000 Ignore all unused code
package main

import "slices"

// 時間計算量： 両者とも最悪で指数時間かかるが、バックトラッキングの方が剪定の効果があり実行時間は短くなりやすい。
// 空間計算量： DPはすべての部分和の組み合わせを記録するためメモリ使用量が多くなる。バックトラッキングは再帰の深さ分だけを保持するためメモリ使用量は少なくなる。

func combinationSumBacktrackingRecursion(candidates []int, target int) [][]int {
	var combinations [][]int
	var combination []int
	var generate func(int, int)
	generate = func(sum, currentIndex int) {
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
			generate(sum+candidates[i], i)
			combination = combination[:len(combination)-1]
		}
	}
	generate(0, 0)
	return combinations
}

type combinationFrame struct {
	combination  []int
	sum          int
	currentIndex int
}

func combinationSum(candidates []int, target int) [][]int {
	var combinations [][]int
	stack := []combinationFrame{{[]int{}, 0, 0}}
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
			combination := slices.Clone(f.combination)
			combination = append(combination, candidates[i])
			stack = append(stack, combinationFrame{combination, f.sum + candidates[i], i})
		}
	}
	return combinations
}

func combinationSumDP(candidates []int, target int) [][]int {
	memo := make([][][]int, target+1)
	memo[0] = [][]int{{}}
	for _, candidate := range candidates {
		for sum := candidate; sum <= target; sum++ {
			for _, combination := range memo[sum-candidate] {
				newCombination := slices.Clone(combination)
				newCombination = append(newCombination, candidate)
				memo[sum] = append(memo[sum], newCombination)
			}
		}
	}
	return memo[target]
}
