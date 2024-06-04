//lint:file-ignore U1000 Ignore all unused code
package main

func romanToInt(s string) int {
	symbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	res := 0
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) &&
			symbols[string(runes[i])] < symbols[string(runes[i+1])] {
			res -= symbols[string(runes[i])]
		} else {
			res += symbols[string(runes[i])]
		}
	}

	return res
}
