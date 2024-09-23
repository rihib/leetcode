//lint:file-ignore U1000 Ignore all unused code
package stringtointegeratoi

import "math"

/*
時間：20分
解法自体はすぐに思いついたのですが、色々とバグを取り除くのに時間がかかってしまいました。
*/
func myAtoi_step1(s string) int {
	const (
		intMax = int(math.MaxInt32)
		intMin = int(math.MinInt32)
	)
	sign := 1
	i := 0
	num := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < len(s) && s[i] == '0' {
		i++
	}
	for i < len(s) && '0' <= rune(s[i]) && rune(s[i]) <= '9' {
		num = num*10 + int(rune(s[i])-'0')
		if sign == 1 && num > intMax {
			return intMax
		}
		if sign == -1 && -num < intMin {
			return intMin
		}
		i++
	}
	return num * sign
}
