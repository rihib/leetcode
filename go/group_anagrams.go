//lint:file-ignore U1000 Ignore all unused code
package main

func groupAnagrams(strs []string) [][]string {
	anagramsMap := make(map[[26]int][]string, len(strs))
	for _, word := range strs {
		var frequencies [26]int
		for _, r := range word {
			frequencies[r-'a']++
		}
		anagramsMap[frequencies] = append(anagramsMap[frequencies], word)
	}
	anagrams := make([][]string, 0, len(anagramsMap))
	for _, group := range anagramsMap {
		anagrams = append(anagrams, group)
	}
	return anagrams
}
