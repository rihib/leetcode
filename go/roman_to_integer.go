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
	currIndex, nextIndex := 0, 1
	for currIndex < len(runeS) && nextIndex < len(runeS) {
		curr, next := runeS[currIndex], runeS[nextIndex]
		if (curr == 'I' && (next == 'V' || next == 'X')) ||
			(curr == 'X' && (next == 'L' || next == 'C')) ||
			(curr == 'C' && (next == 'D' || next == 'M')) {
			total += symbolToInt[next] - symbolToInt[curr]
			currIndex, nextIndex = currIndex+2, nextIndex+2
			continue
		}
		total += symbolToInt[curr]
		currIndex, nextIndex = currIndex+1, nextIndex+1
	}
	if currIndex < len(runeS) {
		total += symbolToInt[runeS[currIndex]]
	}
	return total
}

func romanToInt2(s string) int {
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
