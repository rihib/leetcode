//lint:file-ignore U1000 Ignore all unused code
package main

func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		count += n % 2
		n /= 2
	}
	return count
}
