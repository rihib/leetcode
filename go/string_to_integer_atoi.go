//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"math"
	"strconv"
)

func myAtoi(s string) int {
	currentIndex := 0
	for currentIndex < len(s) && s[currentIndex] == ' ' {
		currentIndex++
	}
	signum := 1
	if currentIndex < len(s) {
		switch s[currentIndex] {
		case '+':
			currentIndex++
		case '-':
			signum = -1
			currentIndex++
		}
	}
	n := 0
	for i := currentIndex; i < len(s); i++ {
		digit, err := strconv.Atoi(string(s[i]))
		if err != nil {
			break
		}
		if signum == 1 && (n > math.MaxInt32/10 || digit >= math.MaxInt32-n*10) {
			return math.MaxInt32
		}
		if signum == -1 && (-n < math.MinInt32/10 || -digit <= math.MinInt32- -n*10) {
			return math.MinInt32
		}
		n = n*10 + digit
	}
	return n * signum
}
