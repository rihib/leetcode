//lint:file-ignore U1000 Ignore all unused code
package longestcommonprefix

import "strings"

/*
レビュワーの方へ：
	- このコードは既にGoの標準のフォーマッタで整形済みです。演算子の周りにスペースがあったりなかったりしますが、これはGoのフォーマッタによるもので、優先順位の高い演算子の周りにはスペースが入らず、低い演算子の周りには入るようになっています。https://qiita.com/tchssk/items/77030b4271cd192d0347
*/

/*
トライ木でも解けるとのことなので、初実装してみた。トライ木といっても単に木なので特に問題なく実装できた。
*/
type Trie struct {
	root *Node
}

type Node struct {
	children  map[rune]*Node
	isWordEnd bool
}

func NewTrie() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			node.children[r] = &Node{children: make(map[rune]*Node)}
		}
		node = node.children[r]
	}
	node.isWordEnd = true
}

func (t *Trie) Prefix() string {
	var prefix strings.Builder
	node := t.root
	for len(node.children) == 1 && !node.isWordEnd {
		for r, child := range node.children {
			prefix.WriteRune(r)
			node = child // ここでnodeをchildで更新してもrangeは評価済みなのでループが続くことはない
		}
	}
	return prefix.String()
}

func longestCommonPrefixTrie(strs []string) string {
	t := NewTrie()
	for _, word := range strs {
		t.Insert(word)
	}
	return t.Prefix()
}
