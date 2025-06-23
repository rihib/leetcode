//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	sign := 1
	if i < len(s) {
		switch s[i] {
		case '+':
			i++
		case '-':
			sign = -1
			i++
		}
	}
	n := 0
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')
		if sign == 1 {
			if n > math.MaxInt32/10 || math.MaxInt32-n*10 <= digit {
				return math.MaxInt32
			}
		} else {
			if -n < math.MinInt32/10 || math.MinInt32- -n*10 >= -digit {
				return math.MinInt32
			}
		}
		n = n*10 + digit
		i++
	}
	return n * sign
}
