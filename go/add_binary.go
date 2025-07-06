//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"slices"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	var reversed strings.Builder
	carry := 0
	for i := 1; ; i++ {
		if len(a)-i < 0 && len(b)-i < 0 && carry == 0 {
			break
		}
		bitA, bitB := 0, 0
		if len(a)-i >= 0 {
			bitA, _ = strconv.Atoi(string(a[len(a)-i]))
		}
		if len(b)-i >= 0 {
			bitB, _ = strconv.Atoi(string(b[len(b)-i]))
		}
		sum := bitA + bitB + carry
		reversed.WriteString(strconv.Itoa(sum % 2))
		carry = sum / 2
	}
	added := []rune(reversed.String())
	slices.Reverse(added)
	return string(added)
}
