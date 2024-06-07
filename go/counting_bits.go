//lint:file-ignore U1000 Ignore all unused code
package main

func countBits(n int) []int {
	res := make([]int, n+1)
	powerOf2 := 1
	for i := 1; i <= n; i++ {
		if powerOf2*2 == i {
			powerOf2 = i
		}
		res[i] = 1 + res[i-powerOf2]
	}
	return res
}
