//lint:file-ignore U1000 Ignore all unused code
package main

func permute_backtracking_iterative(nums []int) [][]int {
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
			stack = append(stack, state{
				permutation: newPermutation,
				inUse:       newInUse,
			})
		}
	}
	return permutations
}

func permute_backtracking_recursion(nums []int) [][]int {
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
