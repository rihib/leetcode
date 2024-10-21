//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"sort"
)

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLength := len(strs[0])
	for _, s := range strs {
		minLength = min(minLength, len(s))
	}
	for i := 0; i < minLength; i++ {
		c := strs[0][i]
		for _, word := range strs[1:] {
			if word[i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLength]
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
