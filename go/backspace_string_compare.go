//lint:file-ignore U1000 Ignore all unused code
package main

func backspaceCompare(s string, t string) bool {
	runeS, runeT := []rune(s), []rune(t)
	i, j := len(runeS)-1, len(runeT)-1
	for i >= 0 || j >= 0 {
		i, j = nextIndex(runeS, i), nextIndex(runeT, j)
		if i >= 0 && j >= 0 && runeS[i] == runeT[j] {
			i--
			j--
			continue
		}
		if i >= 0 || j >= 0 {
			return false
		}
	}
	return true
}

func nextIndex(runeS []rune, index int) int {
	count := 0
	for {
		if index >= 0 && runeS[index] == '#' {
			count++
		} else if count > 0 {
			count--
		} else {
			break
		}
		index--
	}
	return index
}

func backspaceCompare2(s string, t string) bool {
	runeS, runeT := []rune(s), []rune(t)
	i, j := len(runeS)-1, len(runeT)-1
	sCount, tCount := 0, 0
	for i >= 0 || j >= 0 {
		if i >= 0 && runeS[i] == '#' {
			sCount++
			i--
			continue
		}
		if j >= 0 && runeT[j] == '#' {
			tCount++
			j--
			continue
		}
		if sCount > 0 {
			sCount--
			i--
			continue
		}
		if tCount > 0 {
			tCount--
			j--
			continue
		}
		if i >= 0 && j >= 0 && runeS[i] == runeT[j] {
			i--
			j--
			continue
		}
		return false
	}
	return true
}
