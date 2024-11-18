//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var prefix strings.Builder
	for i := 0; ; i++ {
		var curr rune
		for _, word := range strs {
			if i >= len(word) {
				return prefix.String()
			}
			if curr == 0 {
				curr = rune(word[i])
				continue
			}
			if rune(word[i]) != curr {
				return prefix.String()
			}
		}
		prefix.WriteRune(curr)
	}
}

/*
Trie
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
