//lint:file-ignore U1000 Ignore all unused code
package main

import "sort"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)
	for i := range strs[0] {
		if strs[0][i] != strs[len(strs)-1][i] {
			return strs[0][:i]
		}
	}
	return strs[0]
}
