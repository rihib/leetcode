//lint:file-ignore U1000 Ignore all unused code
package main

func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	original, reversed := x, 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed
}
