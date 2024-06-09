//lint:file-ignore U1000 Ignore all unused code
package main

func hammingWeight(n int) int {
	cnt := 0
	for n > 0 {
		n &= n - 1
		cnt += 1
	}
	return cnt
}
