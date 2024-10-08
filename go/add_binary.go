//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"slices"
	"strings"
)

func addBinary(a string, b string) string {
	var reversed strings.Builder
	maxLength := max(len(a), len(b))
	carry := 0
	for i := 1; i <= maxLength; i++ {
		bitA, bitB := 0, 0
		if i <= len(a) {
			bitA = int(a[len(a)-i] - '0') // a[len(a)-i]がbyte型なので'0'もbyte型と解釈される
		}
		if i <= len(b) {
			bitB = int(b[len(b)-i] - '0')
		}
		sum := bitA + bitB + carry
		carry = sum / 2
		reversed.WriteByte(byte(sum%2 + '0')) // sum%2がint型なので'0'もint型として扱われる
	}
	if carry == 1 {
		reversed.WriteByte(byte(carry + '0'))
	}
	result := []rune(reversed.String())
	slices.Reverse(result)
	return string(result)
}
