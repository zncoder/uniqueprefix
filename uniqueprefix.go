// Package uniqueprefix computes the unique prefixes of a list of words.
// It implements with a Trie.
package uniqueprefix

import (
	"fmt"
	"io"
)

// Prefixes computes the unique prefixes of names.
func Prefixes(names ...string) (prefixes []string, unique bool) {
	unique = true

	trie := NewTrie()
	for _, name := range names {
		if !trie.Add(name) {
			unique = false
		}
	}

	for _, name := range names {
		p, ok := trie.Prefix(name)
		if !ok {
			panic(fmt.Sprintf("BUG: get prefix of %q of %v", name, names))
		}
		prefixes = append(prefixes, p)
	}
	return prefixes, unique
}

// Trie implements a trie of words.
type Trie struct {
	root trieNode
}

func NewTrie() *Trie {
	return &Trie{root: trieNode{multi: true}}
}

type trieNode struct {
	r      rune
	multi  bool
	fanout []trieNode
}

func printTrieNode(cur *trieNode, w io.Writer, prefix string) {
	io.WriteString(w, prefix)
	if cur.r != 0 {
		io.WriteString(w, string(cur.r))
	}
	if cur.multi {
		io.WriteString(w, "+")
	}
	io.WriteString(w, "\n")
	for i := range cur.fanout {
		printTrieNode(&cur.fanout[i], w, prefix+"  ")
	}
}

// Add adds a word s to the trie. It returns true if the word does not
// exist in the trie already.
func (tr *Trie) Add(s string) (added bool) {
	cur := &tr.root
L:
	for _, r := range s {
		for i := range cur.fanout {
			if cur.fanout[i].r == r {
				cur = &cur.fanout[i]
				cur.multi = true
				continue L
			}
		}
		cur.fanout = append(cur.fanout, trieNode{r: r})
		cur = &cur.fanout[len(cur.fanout)-1]
		added = true
	}
	return added
}

// Prefix returns the unique prefix of a word s. The word s should be
// added already to the trie, otherwise ok is false.
func (tr *Trie) Prefix(s string) (prefix string, ok bool) {
	cur := &tr.root
L:
	for i, r := range s {
		if !cur.multi && prefix == "" {
			// cur cannot be root, because root.multi is true
			if i == 0 {
				panic(fmt.Sprintf("BUG: cur:%v s:%s", cur, s))
			}
			prefix = s[:i]
			// continue to make sure s is in tr
		}

		for j := range cur.fanout {
			if cur.fanout[j].r == r {
				cur = &cur.fanout[j]
				continue L
			}
		}
		// s is not in trie
		return prefix, false
	}
	// full match
	if prefix == "" {
		prefix = s
	}
	return prefix, true
}
