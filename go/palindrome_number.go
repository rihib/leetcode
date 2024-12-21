//lint:file-ignore U1000 Ignore all unused code
package main

import "strconv"

func isPalindromeNumbver1(x int) bool {
	var reversed int
	tmp := x
	for tmp > 0 {
		n := tmp % 10
		reversed = reversed*10 + n
		tmp /= 10
	}
	return x == reversed
}

// Itoaは数字を文字列に変換するのでエラーが起きることはないため、返り値としてerrorは返らない。
// Atoiは文字列を数字に変換するので変換エラーが起きる可能性があり、返り値としてerrorも返る。
func isPalindromeNumber2(x int) bool {
	strX := strconv.Itoa(x)
	i, j := 0, len(strX)-1
	for i < j {
		if strX[i] != strX[j] {
			return false
		}
		i++
		j--
	}
	return true
}
