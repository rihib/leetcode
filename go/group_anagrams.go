//lint:file-ignore U1000 Ignore all unused code
package main

func groupAnagrams(strs []string) [][]string {
	anagramsMap := make(map[[26]int][]string)
	for _, word := range strs {
		var frequencies [26]int
		for _, r := range word {
			frequencies[r-'a']++
		}
		anagramsMap[frequencies] = append(anagramsMap[frequencies], word)
	}

	anagrams := make([][]string, len(anagramsMap))
	i := 0
	for _, words := range anagramsMap {
		anagrams[i] = words
		i++
	}
	return anagrams
}
