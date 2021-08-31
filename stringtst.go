// Package stringtst implements Ternary Search Trees for strings
package stringtst

type trie struct {
	key                 int
	complete            bool
	left, middle, right *trie
}

// TST is a container for the tree root node
type TST struct {
	root *trie // Root of the tree
}

func isWord(s string) bool {
	for i := range s {
		if int("a"[0]) <= i && i <= int("z"[0]) {
			return false
		}
	}
	return true
}

func insideRange(t *trie, lower, upper int) bool {
	if t == nil {
		return true
	}
	if !(lower < t.key && t.key < upper) {
		return false
	}

	return insideRange(t.left, lower, t.key) && isTrie(t.middle) && insideRange(t.right, t.key, upper)
}

func isTrie(t *trie) bool {
	return insideRange(t, int("a"[0])-1, int("z"[0])+1)
}

func isTST(t *TST) bool {
	return isTrie(t.root)
}

// NewTST returns a new TST with an empty tree
func NewTST() *TST {
	return &TST{root: nil}
}

// TrieSearch searches for string s in trie t
// (current char at pos i in s)
func trieSearch(t *trie, s string, i int) bool {
	if t == nil {
		return false
	}

	c := int(s[i])

	if c == t.key {
		// Current char matches current node's char
		if i == len(s)-1 {
			return t.complete // s is a complete word in t
		}
		return trieSearch(t.middle, s, i+1)
	} else if c < t.key {
		return trieSearch(t.left, s, i)
	} else {
		return trieSearch(t.right, s, i)
	}
}

//TSTSearch searches for string s in trie t
func TSTSearch(t *TST, s string) bool {
	return trieSearch(t.root, s, 0)
}

func trieInsert(t *trie, s string, i int) *trie {
	c := int(s[i])

	if t == nil {
		// Node is empty, init a new node
		t = &trie{key: c, complete: false, left: nil, middle: nil, right: nil}
	}

	if c == t.key {
		if i == len(s)-1 {
			t.complete = true // End of string
			return t
		}
		// Still chars left, recursively insert rest of string
		t.middle = trieInsert(t.middle, s, i+1)
	} else if c < t.key {
		t.left = trieInsert(t.left, s, i)
	} else {
		t.right = trieInsert(t.right, s, i)
	}

	return t
}

// TSTInsert inserts string s in TST t
func TSTInsert(t *TST, s string) {
	t.root = trieInsert(t.root, s, 0)
}

func trieHasPrefix(t *trie, s string) bool {
	if len(s) == 0 {
		return true // Empty string, prefix exists
	} else if t == nil {
		return false // Empty node, string is not empty, prefix not in t
	} else {
		// Keep checking next chars
		c := int(s[0])
		if t.key == c {
			return trieHasPrefix(t.middle, s[1:])
		} else if t.key < c {
			return trieHasPrefix(t.right, s[1:])
		} else {
			return trieHasPrefix(t.left, s[1:])
		}
	}
}

// TSTHasPrefix checks if prefix s is in t
func TSTHasPrefix(t *TST, s string) bool {
	return trieHasPrefix(t.root, s)
}
