//lint:file-ignore U1000 Ignore all unused code
package main

func countBits(n int) []int {
	bitCounts := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bitCounts[i] = 1 + bitCounts[i&(i-1)] // i&(i-1) unset rightmost set bit
	}
	return bitCounts
}
