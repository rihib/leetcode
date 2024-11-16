//lint:file-ignore U1000 Ignore all unused code
package main

func countBits(n int) []int {
	bitCounts := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			bitCounts[i] = i
			continue
		}
		j := i & (i - 1)
		bitCounts[i] = bitCounts[j] + 1
	}
	return bitCounts
}
