//lint:file-ignore U1000 Ignore all unused code
package main

func hammingWeightBitmanipulation(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}

func hammingWeight2Complement(n int) int {
	count := 0
	for n > 0 {
		n -= n & -n
		count++
	}
	return count
}
