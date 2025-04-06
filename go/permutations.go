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
	var permutations [][]int
	permutation := slices.Clone(nums)
	sort.Ints(permutation)
	for {
		newPermutation := slices.Clone(permutation)
		permutations = append(permutations, newPermutation)
		sortedUntil := len(permutation) - 2
		for sortedUntil >= 0 && permutation[sortedUntil] > permutation[sortedUntil+1] {
			sortedUntil--
		}
		if sortedUntil < 0 {
			break
		}
		swapTarget := len(permutation) - 1
		for permutation[sortedUntil] > permutation[swapTarget] {
			swapTarget--
		}
		permutation[sortedUntil], permutation[swapTarget] = permutation[swapTarget], permutation[sortedUntil]
		sort.Ints(permutation[sortedUntil+1:])
	}
	return permutations
}

func permuteBacktrackingRecursion(nums []int) [][]int {
	var permutations [][]int
	permutation := make([]int, 0, len(nums))
	inUse := make(map[int]struct{}, len(nums))
	var generate func()
	generate = func() {
		if len(inUse) == len(nums) {
			newPermutation := slices.Clone(permutation)
			permutations = append(permutations, newPermutation)
			return
		}
		for _, n := range nums {
			if _, ok := inUse[n]; ok {
				continue
			}
			permutation = append(permutation, n)
			inUse[n] = struct{}{}
			generate()
			permutation = permutation[:len(permutation)-1]
			delete(inUse, n)
		}
	}
	generate()
	return permutations
}

type permutationFrame struct {
	permutation []int
	inUse       map[int]struct{}
}

func permuteBacktrackingIterative(nums []int) [][]int {
	var permutations [][]int
	permutation := make([]int, 0, len(nums))
	inUse := make(map[int]struct{}, len(nums))
	stack := []permutationFrame{{permutation, inUse}}
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(f.permutation) == len(nums) {
			permutations = append(permutations, f.permutation)
			continue
		}
		for _, n := range nums {
			if _, ok := f.inUse[n]; ok {
				continue
			}
			newPermutation := slices.Clone(f.permutation)
			newPermutation = append(newPermutation, n)
			newInUse := maps.Clone(f.inUse)
			newInUse[n] = struct{}{}
			stack = append(stack, permutationFrame{newPermutation, newInUse})
		}
	}
	return permutations
}
