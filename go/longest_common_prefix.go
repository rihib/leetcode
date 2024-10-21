//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"sort"
)

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix []rune
	for i := 0; ; i++ {
		if i >= len(strs[0]) {
			return string(prefix)
		}
		curr := strs[0][i]
		for _, word := range strs {
			if i >= len(word) {
				return string(prefix)
			}
			if word[i] != curr {
				return string(prefix)
			}
		}
		prefix = append(prefix, rune(curr))
	}
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	for i := range strs[0] {
		if strs[0][i] != strs[len(strs)-1][i] {
			return strs[0][:i]
		}
	}
	return strs[0]
}
