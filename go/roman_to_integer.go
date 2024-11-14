//lint:file-ignore U1000 Ignore all unused code
package main

func romanToInt(s string) int {
	symbolToInt := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	runeS := []rune(s)
	for i, r := range runeS {
		if i+1 < len(runeS) && symbolToInt[r] < symbolToInt[runeS[i+1]] {
			total -= symbolToInt[r]
		} else {
			total += symbolToInt[r]
		}
	}
	return total
}
