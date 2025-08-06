//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"math"
	"strconv"
)

func myAtoi(s string) int {
	index := 0
	for index < len(s) && s[index] == ' ' {
		index++
	}
	signum := 1
	if index < len(s) {
		switch s[index] {
		case '+':
			index++
		case '-':
			signum = -1
			index++
		}
	}
	n := 0
	for i := index; i < len(s); i++ {
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
