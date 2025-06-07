//lint:file-ignore U1000 Ignore all unused code
package main

import "strings"

func backspaceCompareStack(s string, t string) bool {
	return typedText(s) == typedText(t)
}

func typedText(s string) string {
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
	return reversedText(s) == reversedText(t)
}

func reversedText(s string) string {
	var text strings.Builder
	skip := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			skip++
		} else if skip > 0 {
			skip--
		} else {
			text.WriteByte(s[i])
		}
	}
	return text.String()
}
