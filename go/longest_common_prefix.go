//lint:file-ignore U1000 Ignore all unused code
package main

import "strings"

func longestCommonPrefix(strs []string) string {
	var prefix strings.Builder
	for i := 0; i < len(strs[0]); i++ {
		for _, word := range strs {
			if i == len(word) || strs[0][i] != word[i] {
				return prefix.String()
			}
		}
		prefix.WriteByte(strs[0][i])
	}
	return prefix.String()
}
