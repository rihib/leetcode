//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"slices"
	"strings"
)

func backspaceCompareStack(s string, t string) bool {
	return typedText1(s) == typedText1(t)
}

func typedText1(s string) string {
	stack := make([]rune, 0, len(s))
	for _, r := range s {
		if r == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, r)
	}
	return string(stack)
}

func backspaceCompareReverse(s string, t string) bool {
	return typedText2(s) == typedText2(t)
}

func typedText2(s string) string {
	var text strings.Builder
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			count++
		} else if count > 0 {
			count--
		} else {
			text.WriteByte(s[i])
		}
	}
	runes := []rune(text.String())
	slices.Reverse(runes)
	return string(runes)
}
