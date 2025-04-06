//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"maps"
	"slices"
	"sort"
)

func permuteLexicographically(nums []int) [][]int {
	sort.Ints(nums)
	var permutations [][]int
	for {
		permutations = append(permutations, append([]int{}, nums...))
		i := len(nums) - 2
		for i >= 0 && nums[i] >= nums[i+1] {
			i--
		}
		if i < 0 {
			break
		}
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		reverse(nums[i+1:])
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
