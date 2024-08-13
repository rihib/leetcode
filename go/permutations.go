//lint:file-ignore U1000 Ignore all unused code
package main

func permute(nums []int) [][]int {
	var permutations [][]int
	stack := make([]int, 0, len(nums))
	inUse := make(map[int]struct{})
	var generate func()
	generate = func() {
		if len(stack) == len(nums) {
			permutations = append(permutations, append([]int{}, stack...))
			return
		}
		for _, n := range nums {
			if _, ok := inUse[n]; !ok {
				stack = append(stack, n)
				inUse[n] = struct{}{}
				generate()
				stack = stack[:len(stack)-1]
				delete(inUse, n)
			}
		}
	}
	generate()
	return permutations
}
