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
	combination := slices.Clone(nums)
	sort.Ints(combination)
	for {
		combinations = append(combinations, slices.Clone(combination))
		sortedUntil := len(combination) - 2
		for sortedUntil >= 0 && combination[sortedUntil] > combination[sortedUntil+1] {
			sortedUntil--
		}
		if sortedUntil < 0 {
			break
		}
		target := len(combination) - 1
		for combination[sortedUntil] >= combination[target] {
			target--
		}
		combination[sortedUntil], combination[target] = combination[target], combination[sortedUntil]
		sort.Ints(combination[sortedUntil+1:])
	}
	return combinations
}

func permuteBacktrackingRecursion(nums []int) [][]int {
	var combinations [][]int
	combination := make([]int, 0, len(nums))
	seen := make(map[int]struct{}, len(nums))
	var generate func()
	generate = func() {
		if len(combination) == len(nums) {
			combinations = append(combinations, slices.Clone(combination))
			return
		}
		for _, n := range nums {
			if _, ok := seen[n]; ok {
				continue
			}
			combination = append(combination, n)
			seen[n] = struct{}{}
			generate()
			combination = combination[:len(combination)-1]
			delete(seen, n)
		}
	}
	generate()
	return combinations
}

type permutationFrame struct {
	combination []int
	seen        map[int]struct{}
}

func permuteBacktrackingIterative(nums []int) [][]int {
	var combinations [][]int
	stack := []permutationFrame{{make([]int, 0, len(nums)), make(map[int]struct{}, len(nums))}}
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(f.combination) == len(nums) {
			combinations = append(combinations, f.combination)
			continue
		}
		for _, n := range nums {
			if _, ok := f.seen[n]; ok {
				continue
			}
			newCombination := append(slices.Clone(f.combination), n)
			newSeen := maps.Clone(f.seen)
			newSeen[n] = struct{}{}
			stack = append(stack, permutationFrame{newCombination, newSeen})
		}
	}
	return combinations
}
