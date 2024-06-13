//lint:file-ignore U1000 Ignore all unused code
package main

func climbStairs(n int) int {
	prev, curr := 0, 1
	for i := 0; i < n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
