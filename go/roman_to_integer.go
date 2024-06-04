//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"log"
)

func romanToInt(s string) int {
	s_list := []rune(s)
	result := 0

	symbols := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		n, ok := symbols[s_list[i]]

		if !ok {
			log.Fatal("Invalid symbol")
		}

		if i+1 < len(s) {
			if s_list[i] == 'I' {
				if s_list[i+1] == 'V' {
					result += 4
					i++
					continue
				}
				if s_list[i+1] == 'X' {
					result += 9
					i++
					continue
				}
			}
			if s_list[i] == 'X' {
				if s_list[i+1] == 'L' {
					result += 40
					i++
					continue
				}
				if s_list[i+1] == 'C' {
					result += 90
					i++
					continue
				}
			}
			if s_list[i] == 'C' {
				if s_list[i+1] == 'D' {
					result += 400
					i++
					continue
				}
				if s_list[i+1] == 'M' {
					result += 900
					i++
					continue
				}
			}
		}

		result += n
	}

	return result
}
