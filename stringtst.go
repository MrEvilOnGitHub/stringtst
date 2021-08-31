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

	return insideRange(t.left, lower, t.key) && t.middle.isTrie() && insideRange(t.right, t.key, upper)
}

func (t *trie) isTrie() bool {
	return insideRange(t, int("a"[0])-1, int("z"[0])+1)
}

func (t *TST) isTST() bool {
	return t.root.isTrie()
}

// NewTST returns a new TST with an empty tree
func NewTST() *TST {
	return &TST{root: nil}
}

// TrieSearch searches for string s in trie t
// (current char at pos i in s)
func (t *trie) search(s string, i int) bool {
	if t == nil {
		return false
	}

	c := int(s[i])

	if c == t.key {
		// Current char matches current node's char
		if i == len(s)-1 {
			return t.complete // s is a complete word in t
		}
		return t.middle.search(s, i+1)
	} else if c < t.key {
		return t.left.search(s, i)
	} else {
		return t.right.search(s, i)
	}
}

// Search searches for string s in trie t
func (t *TST) Search(s string) bool {
	return t.root.search(s, 0)
}

func (t *trie) insert(s string, i int) *trie {
  if i >= len(s) {
    return t
  }


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
		t.middle = t.middle.insert(s, i+1)
	} else if c < t.key {
		t.left = t.left.insert(s, i)
	} else {
		t.right = t.right.insert(s, i)
	}

	return t
}

// Insert inserts string s in TST t
func (t *TST) Insert(s string) {
	t.root = t.root.insert(s, 0)
}

func (t *trie) hasPrefix(s string) bool {
	if len(s) == 0 {
		return true // Empty string, prefix exists
	} else if t == nil {
		return false // Empty node, string is not empty, prefix not in t
	} else {
		// Keep checking next chars
		c := int(s[0])
		if t.key == c {
			return t.middle.hasPrefix(s[1:])
		} else if t.key < c {
			return t.right.hasPrefix(s[1:])
		} else {
			return t.left.hasPrefix(s[1:])
		}
	}
}

// HasPrefix checks if prefix s is in t
func (t *TST) HasPrefix(s string) bool {
	return t.root.hasPrefix(s)
}
