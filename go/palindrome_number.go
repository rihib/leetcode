//lint:file-ignore U1000 Ignore all unused code
package main

func isPalindromeNumbver(x int) bool {
	reversed := 0
	tmp := x
	for tmp > 0 {
		reversed = reversed*10 + tmp%10
		tmp /= 10
	}
	return x == reversed
}
