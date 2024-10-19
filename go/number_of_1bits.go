//lint:file-ignore U1000 Ignore all unused code
package main

func hammingWeightNaive1(n int) int {
	count := 0
	for n > 0 {
		count += n % 2
		n /= 2
	}
	return count
}

func hammingWeightNaive2(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func hammingWeightBitmanipulation(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count += 1
	}
	return count
}

func hammingWeight2Complement(n int) int {
	count := 0
	for n > 0 {
		n -= (n & -n)
		count += 1
	}
	return count
}
