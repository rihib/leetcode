//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"sort"
	"strings"
)

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLength := len(strs[0])
	for _, s := range strs {
		minLength = min(minLength, len(s))
	}
	for i := 0; i < minLength; i++ {
		c := strs[0][i]
		for _, word := range strs[1:] {
			if word[i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLength]
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	for i := range strs[0] {
		if strs[0][i] != strs[len(strs)-1][i] {
			return strs[0][:i]
		}
	}
	return strs[0]
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

func longestCommonPrefix(strs []string) string {
	t := NewTrie()
	for _, word := range strs {
		t.Insert(word)
	}
	return t.Prefix()
}
