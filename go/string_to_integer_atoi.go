//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	runeS := []rune(s)
	currentIndex := 0
	for currentIndex < len(runeS) && runeS[currentIndex] == ' ' {
		currentIndex++
	}
	sign := 1
	if currentIndex < len(runeS) && (runeS[currentIndex] == '+' || runeS[currentIndex] == '-') {
		if runeS[currentIndex] == '-' {
			sign = -1
		}
		currentIndex++
	}
	num := 0
	for currentIndex < len(runeS) && unicode.IsDigit(runeS[currentIndex]) {
		digit := int(runeS[currentIndex] - '0')
		if sign == 1 && (math.MaxInt32/10 < num || math.MaxInt32/10 == num && math.MaxInt32%10 < digit) {
			return math.MaxInt32
		}
		if sign == -1 && (math.MinInt32/10 > -num || math.MinInt32/10 == -num && math.MinInt32%10 > -digit) {
			return math.MinInt32
		}
		num = num*10 + digit
		currentIndex++
	}
	return sign * num
}
