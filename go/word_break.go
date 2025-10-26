//lint:file-ignore U1000 Ignore all unused code
package main

// DPを使うと線型時間で解ける
func wordBreak(s string, wordDict []string) bool {
	canBreak := make([]bool, len(s)+1)
	canBreak[0] = true
	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			if i >= len(word) && canBreak[i-len(word)] && s[i-len(word):i] == word {
				canBreak[i] = true
				break // この箇所がすでに分割可能と分かれば他の単語でも分割可能かを調べる必要はない
			}
		}
	}
	return canBreak[len(s)]
}

// 下記のようにTrieを使った実装も可能だが、下記のような場合は計算量が指数時間になってしまう。
// "aaa..."という部分文字列はどの単語でもマッチするため、再帰呼び出しが爆発的に増える。
// sの各位置で分割するかしないかの2択になり、2^n通りの分割方法を試すことになる。
// しかも下記は最後にbがあるため、最終的に分割不可能と判定されるまで全通りを試すことになる。
// s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
// wordDict = ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa"]
func wordBreakBad(s string, wordDict []string) bool {
	root := trie(wordDict)
	return canBeSegmented(s, root)
}

func canBeSegmented(s string, root *trieNodeBad) bool {
	node := root
	for i, r := range s {
		next, ok := node.children[r]
		if !ok {
			break
		}
		if next.isWordEnd && i == len(s)-1 ||
			next.isWordEnd && i+1 < len(s) && canBeSegmented(s[i+1:], root) {
			return true
		}
		node = next
	}
	return false
}

type trieNodeBad struct {
	children  map[rune]*trieNodeBad
	isWordEnd bool
}

func trie(wordDict []string) *trieNodeBad {
	root := &trieNodeBad{make(map[rune]*trieNodeBad), false}
	node := root
	for _, word := range wordDict {
		for _, r := range word {
			if _, ok := node.children[r]; !ok {
				node.children[r] = &trieNodeBad{make(map[rune]*trieNodeBad), false}
			}
			node = node.children[r]
		}
		node.isWordEnd = true
		node = root
	}
	return root
}
