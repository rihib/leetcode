//lint:file-ignore U1000 Ignore all unused code
package permutations

/*
時間：10分
いくつか以前にバックトラッキングの問題を解いたことがあるので割とすぐに書くことができました。
*/
func permuteStep1(nums []int) [][]int {
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
	inUse := make(map[int]struct{})
	generate(inUse)
	return permutations
}
