//lint:file-ignore U1000 Ignore all unused code
package main

import "math"

func myAtoi(s string) int {
	const (
		intMax = int(math.MaxInt32)
		intMin = int(math.MinInt32)
	)

	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	sign := 1
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	num := 0
	for i < len(s) && '0' <= s[i] && s[i] <= '9' {
		next := int(s[i] - '0')
		if sign != -1 && num > (intMax-next)/10 {
			return intMax
		}
		if sign == -1 && -num < (intMin+next)/10 {
			return intMin
		}
		num = num*10 + next
		i++
	}
	return sign * num
}
