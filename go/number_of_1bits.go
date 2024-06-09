//lint:file-ignore U1000 Ignore all unused code
package main

func hammingWeight(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n % 2
		n /= 2
	}
	return cnt
}
