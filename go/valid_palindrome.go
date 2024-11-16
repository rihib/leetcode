//lint:file-ignore U1000 Ignore all unused code
package main

import "unicode"

func isPalindrome(s string) bool {
	runeS := []rune(s)
	i, j := 0, len(s)-1
	for i < j {
		if !unicode.IsDigit(runeS[i]) && !unicode.IsLetter(runeS[i]) {
			i++
			continue
		}
		if !unicode.IsDigit(runeS[j]) && !unicode.IsLetter(runeS[j]) {
			j--
			continue
		}
		if unicode.ToLower(runeS[i]) != unicode.ToLower(runeS[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome2(s string) bool {
	left := make(chan rune)
	right := make(chan rune)
	go filter(s, left, false)
	go filter(s, right, true)
	for {
		l, lok := <-left
		r, rok := <-right
		if !lok || !rok {
			return true
		}
		if l != r {
			return false
		}
	}
}

func filter(s string, c chan rune, reversed bool) {
	for i := range s {
		if reversed {
			i = len(s) - 1 - i
		}
		r := rune(s[i])
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			c <- unicode.ToLower(r)
		}
	}
	close(c)
}
