//lint:file-ignore U1000 Ignore all unused code
package main

import "sort"

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
	sort.Ints(candidates)
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
				break
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

// https://discord.com/channels/1084280443945353267/1233603535862628432/1237744279363915878
// https://github.com/Exzrgs/LeetCode/pull/13#discussion_r1606819961
// https://github.com/fhiyo/leetcode/pull/52#issuecomment-2248269934
// Goは、ランタイムのスタックサイズがデフォルトが小さく、かつスタックサイズの大きさを最小にするような最適化を行っていないため、深さの数だけスタックを浪費する実装になっており、再帰の深さが極端に処理性能が落ちる要因になる。そのため、Goに慣れた人は自然と深い再帰処理はループ処理に書き換えるらしい？
// https://ymotongpoo.hatenablog.com/entry/2015/02/23/165341
// https://zenn.dev/nobonobo/articles/e651c66a15aaed657d6e#%E6%80%A7%E8%83%BD%E3%81%AFgo%E4%B8%AD%E7%B4%9A%E3%81%A7c%2B%2B%E7%8E%84%E4%BA%BA%E3%81%AE9%E5%89%B2%E4%BB%A5%E4%B8%8A
// なぜbacktracking_stackの場合はsortしないとバグる？
// 末尾再帰？
// この前解いたバックトラッキングの問題をDPで解いてみる
// 全ての組み合わせを出すのはバックトラッキングとDP以外になんかあったけ？それぞれのPros Consは？
// appendは容量が余ってるとそのまま要素を入れるとかGPTが言ってなかったっけ

// https://github.com/hayashi-ay/leetcode/pull/4
// https://github.com/hayashi-ay/leetcode/pull/65
// https://github.com/shining-ai/leetcode/pull/52
// https://github.com/SuperHotDogCat/coding-interview/pull/11
// https://github.com/goto-untrapped/Arai60/pull/15
// https://github.com/Exzrgs/LeetCode/pull/13
// https://github.com/nittoco/leetcode/pull/25
// https://github.com/fhiyo/leetcode/pull/52
