//lint:file-ignore U1000 Ignore all unused code
package permutations

import "sort"

/*
他の人のコードを見て、辞書順に順列を生成していくことでもかけることに気がつきました。
基本的にはNext Permutationと同じコードです。
*/
func permuteLexicographicallyStep3(nums []int) [][]int {
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

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func permuteBacktrackingIterativeStep3(nums []int) [][]int {
	var permutations [][]int
	type state struct {
		permutation []int
		inUse       map[int]struct{}
	}
	stack := []state{{[]int{}, make(map[int]struct{})}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(current.permutation) == len(nums) {
			permutations = append(permutations, current.permutation)
			continue
		}
		for _, n := range nums {
			if _, ok := current.inUse[n]; ok {
				continue
			}
			newPermutation := append([]int{}, current.permutation...)
			newPermutation = append(newPermutation, n)
			newInUse := make(map[int]struct{})
			for k, v := range current.inUse {
				newInUse[k] = v
			}
			newInUse[n] = struct{}{}
			stack = append(stack, state{newPermutation, newInUse})
		}
	}
	return permutations
}

func permuteBacktrackingRecursionStep3(nums []int) [][]int {
	var permutations [][]int
	permutation := make([]int, 0, len(nums))
	inUse := make(map[int]struct{})
	var generate func()
	generate = func() {
		if len(permutation) == len(nums) {
			permutations = append(permutations, append([]int{}, permutation...))
			return
		}
		for _, n := range nums {
			if _, ok := inUse[n]; !ok {
				permutation = append(permutation, n)
				inUse[n] = struct{}{}
				generate()
				permutation = permutation[:len(permutation)-1]
				delete(inUse, n)
			}
		}
	}
	generate()
	return permutations
}
