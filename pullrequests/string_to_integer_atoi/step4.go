//lint:file-ignore U1000 Ignore all unused code
package stringtointegeratoi

import "math"

/*
Step3のオーバーフローしてしまうバグを直しました。
*/
func myAtoi_step4(s string) int {
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
		digit := int(s[i] - '0')
		if sign == 1 && (num > intMax/10 || num == intMax/10 && digit >= intMax%10) {
			return intMax
		}
		if sign == -1 && (-num < intMin/10 || -num == intMin/10 && -digit <= intMin%10) {
			return intMin
		}
		num = num*10 + digit
		i++
	}
	return sign * num
}
