//lint:file-ignore U1000 Ignore all unused code
package main

import "slices"

// 時間計算量： 両者とも最悪で指数時間かかるが、バックトラッキングの方が剪定の効果があり実行時間は短くなりやすい。
// 空間計算量： DPはすべての部分和の組み合わせを記録するためメモリ使用量が多くなる。バックトラッキングは再帰の深さ分だけを保持するためメモリ使用量は少なくなる。

func combinationSumBacktrackingRecursion(candidates []int, target int) [][]int {
	var combinations [][]int
	combination := make([]int, 0, len(candidates))
	var generate func(int, int)
	generate = func(sum, currentIndex int) {
		if sum > target {
			return
		}
		if sum == target {
			combinations = append(combinations, slices.Clone(combination))
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

func combinationSumBacktrackingIterative(candidates []int, target int) [][]int {
	var combinations [][]int
	stack := []combinationFrame{{make([]int, 0, len(candidates)), 0, 0}}
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if f.sum > target {
			continue
		}
		if f.sum == target {
			combinations = append(combinations, f.combination)
			continue
		}
		for i := f.currentIndex; i < len(candidates); i++ {
			newCombination := append(slices.Clone(f.combination), candidates[i])
			stack = append(stack, combinationFrame{newCombination, f.sum + candidates[i], i})
		}
	}
	return combinations
}

func combinationSumDP(candidates []int, target int) [][]int {
	memo := make([][][]int, target+1)
	memo[0] = [][]int{{}}
	for _, n := range candidates {
		for sum, combinations := range memo {
			if sum+n > target {
				continue
			}
			for _, comb := range combinations {
				newCombination := append(slices.Clone(comb), n)
				memo[sum+n] = append(memo[sum+n], newCombination)
			}
		}
	}
	return memo[target]
}
