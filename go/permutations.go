//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"maps"
	"slices"
	"sort"
)

// 辞書順に生成していく
// [1, 2, 4, 3]の次は[1, 3, 2, 4]
// まず後ろから昇順になっている箇所を探す（2）
// そしてまた後ろから2までで、2よりも大きい一番小さい数を探す（3）
// そしてその２つを入れ替える（[1, 3, 4, 2]）
// そしてその後ろの部分を逆順にする（[1, 3, 2, 4]）
func permuteLexicographically(nums []int) [][]int {
	var combinations [][]int
	comb := slices.Clone(nums)
	sort.Ints(comb)
	for {
		combinations = append(combinations, slices.Clone(comb))
		sortedUntil := len(comb) - 2
		for sortedUntil >= 0 && comb[sortedUntil] >= comb[sortedUntil+1] {
			sortedUntil--
		}
		if sortedUntil < 0 {
			break
		}
		swapTarget := len(comb) - 1
		for comb[sortedUntil] >= comb[swapTarget] {
			swapTarget--
		}
		comb[sortedUntil], comb[swapTarget] = comb[swapTarget], comb[sortedUntil]
		sort.Ints(comb[sortedUntil+1:])
	}
	return combinations
}

func permuteBacktrackingRecursion(nums []int) [][]int {
	var combinations [][]int
	comb := make([]int, 0, len(nums))
	used := make(map[int]struct{}, len(nums))
	var generate func()
	generate = func() {
		if len(comb) == len(nums) {
			combinations = append(combinations, slices.Clone(comb))
			return
		}
		for _, n := range nums {
			if _, ok := used[n]; ok {
				continue
			}
			comb = append(comb, n)
			used[n] = struct{}{}
			generate()
			comb = comb[:len(comb)-1]
			delete(used, n)
		}
	}
	generate()
	return combinations
}

type permutationFrame struct {
	comb []int
	used map[int]struct{}
}

func permuteBacktrackingIterative(nums []int) [][]int {
	var combinations [][]int
	stack := []permutationFrame{{make([]int, 0, len(nums)), make(map[int]struct{}, len(nums))}}
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(f.comb) == len(nums) {
			combinations = append(combinations, f.comb)
			continue
		}
		for _, n := range nums {
			if _, ok := f.used[n]; ok {
				continue
			}
			newComb := append(slices.Clone(f.comb), n)
			newUsed := maps.Clone(f.used)
			newUsed[n] = struct{}{}
			stack = append(stack, permutationFrame{newComb, newUsed})
		}
	}
	return combinations
}
