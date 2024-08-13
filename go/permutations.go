//lint:file-ignore U1000 Ignore all unused code
package main

func permute(nums []int) [][]int {
	var permutations [][]int
	var stack []int
	var generate func(map[int]struct{})
	generate = func(inUse map[int]struct{}) {
		if len(stack) == len(nums) {
			permutations = append(permutations, append([]int{}, stack...))
			return
		}
		for _, n := range nums {
			if _, ok := inUse[n]; !ok {
				stack = append(stack, n)
				inUse[n] = struct{}{}
				generate(inUse)
				stack = stack[:len(stack)-1]
				delete(inUse, n)
			}
		}
	}
	generate(make(map[int]struct{}))
	return permutations
}
