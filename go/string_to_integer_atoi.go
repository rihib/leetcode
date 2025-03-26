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
	isNegative := false
	if currentIndex < len(runeS) && (runeS[currentIndex] == '+' || runeS[currentIndex] == '-') {
		if runeS[currentIndex] == '-' {
			isNegative = true
		}
		currentIndex++
	}
	num := 0
	for currentIndex < len(runeS) && unicode.IsDigit(runeS[currentIndex]) {
		digit := int(runeS[currentIndex] - '0')
		if !isNegative && (num > math.MaxInt32/10 || num == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			return math.MaxInt32
		}
		if isNegative && (-num < math.MinInt32/10 || -num == math.MinInt32/10 && -digit < math.MinInt32%10) {
			return math.MinInt32
		}
		num = num*10 + digit
		currentIndex++
	}
	if isNegative {
		return -num
	}
	return num
}
