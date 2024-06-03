//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"slices"
	"strings"
)

func addBinary(a, b string) string {
	// TestCase: [0, 1], [1, 1], [11, 1]
	maxLen := max(len(a), len(b))
	var reversed_result strings.Builder
	carry := 0

	for i := 1; i <= maxLen; i++ {
		a_bit, b_bit := 0, 0

		if len(a) >= i {
			a_bit = int(a[len(a)-i] - '0')
		}
		if len(b) >= i {
			b_bit = int(b[len(b)-i] - '0')
		}

		sum := a_bit + b_bit + carry
		carry = sum / 2
		reversed_result.WriteByte(byte(sum%2) + '0')
	}
	if carry > 0 {
		reversed_result.WriteByte(byte(carry) + '0')
	}

	result := []rune(reversed_result.String())
	slices.Reverse(result)
	return string(result)
}
