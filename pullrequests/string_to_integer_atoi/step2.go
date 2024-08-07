//lint:file-ignore U1000 Ignore all unused code
package stringtointegeratoi

import "math"

/*
Goではrune型はint32、byte型はint8のエイリアスで、明示的にキャストしなくても内部でキャストしてくれるので、そのまま比較できるため、rune型にキャストするのをやめました。
*/
func myAtoi_step2(s string) int {
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
		num = num*10 + int(s[i]-'0')
		if sign == 1 && num > intMax {
			return intMax
		}
		if sign == -1 && -num < intMin {
			return intMin
		}
		i++
	}
	return sign * num
}
