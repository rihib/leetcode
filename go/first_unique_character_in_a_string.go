//lint:file-ignore U1000 Ignore all unused code
package main

func firstUniqChar(s string) int {
	frequency := make(map[rune]int)
	for _, r := range s {
		frequency[r]++
	}
	for i, r := range s {
		if frequency[r] == 1 {
			return i
		}
	}
	return -1
}
