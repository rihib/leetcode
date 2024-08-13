//lint:file-ignore U1000 Ignore all unused code
package main

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
