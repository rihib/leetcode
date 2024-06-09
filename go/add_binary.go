//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"slices"
	"strings"
)

func addBinary(a, b string) string {
	// TestCase: [0, 1], [1, 1], [11, 1]
	var reversedRes strings.Builder
	maxLen := max(len(a), len(b))
	carry := 0

	for i := 1; i <= maxLen; i++ {
		bitA, bitB := 0, 0

		if len(a) >= i {
			bitA = int(a[len(a)-i] - '0')
		}
		if len(b) >= i {
			bitB = int(b[len(b)-i] - '0')
		}

		sum := bitA + bitB + carry
		carry = sum / 2
		reversedRes.WriteByte(byte(sum%2) + '0')
	}
	if carry > 0 {
		reversedRes.WriteByte(byte(carry) + '0')
	}

	res := []rune(reversedRes.String())
	slices.Reverse(res)
	return string(res)
}
