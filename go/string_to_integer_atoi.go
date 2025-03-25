//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	var num int
	isNegative := false
	currentIndex := 0
	runeS := []rune(s)
	for currentIndex < len(runeS) && runeS[currentIndex] == ' ' {
		currentIndex++
	}
	if currentIndex < len(runeS) && (runeS[currentIndex] == '+' || runeS[currentIndex] == '-') {
		if runeS[currentIndex] == '-' {
			isNegative = true
		}
		currentIndex++
	}
	for i := currentIndex; i < len(runeS); i++ {
		if !unicode.IsDigit(runeS[i]) {
			break
		}
		digit := int(runeS[i] - '0')
		if !isNegative && (num > math.MaxInt32/10 || num == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			return math.MaxInt32
		}
		if isNegative && (-num < math.MinInt32/10 || -num == math.MinInt32/10 && -digit < math.MinInt32%10) {
			return math.MinInt32
		}
		num = num*10 + digit
	}
	if isNegative {
		return num * -1
	}
	return num
}
