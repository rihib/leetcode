//lint:file-ignore U1000 Ignore all unused code
package main

func romanToInt(s string) int {
	symbols := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	result := 0
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if i+1 < len(runes) && symbols[string(runes[i:i+2])] != 0 {
			result += symbols[string(runes[i:i+2])]
			i++
			continue
		}

		if symbols[string(runes[i])] != 0 {
			result += symbols[string(runes[i])]
		}
	}

	return result
}
