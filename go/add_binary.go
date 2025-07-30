//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"slices"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	var added strings.Builder
	carry := 0
	for i := 1; i <= len(a) || i <= len(b) || carry != 0; i++ {
		bitA, bitB := 0, 0
		if i <= len(a) {
			bitA, _ = strconv.Atoi(string(a[len(a)-i]))
		}
		if i <= len(b) {
			bitB, _ = strconv.Atoi(string(b[len(b)-i]))
		}
		sum := bitA + bitB + carry
		carry = sum / 2
		added.WriteString(strconv.Itoa(sum % 2))
	}
	runes := []rune(added.String())
	slices.Reverse(runes)
	return string(runes)
}
