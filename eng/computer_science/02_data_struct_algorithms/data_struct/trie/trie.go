package trie

/*
Trie (Prefix Tree)
Trie (pronounced as "try") is a tree where each node stores a character, and the path from the root to a node forms a prefix of a string. It is used for efficient storage and searching of strings, especially those with common prefixes.

Meaning and Application

Meaning: Fast search of strings by prefix or full match, memory saving by storing common prefixes.

When to use:
- Autocmplete (e.g., in search engines).
- Spell checking.
- Dictionary storage.
- Compression algorithms (e.g., in IP routing).
- Searching for substrings or prefixes.

### Complexity

| Operation | Time Complexity (O) | Space Complexity (O) |
|:---|:---:|:---:|
| Insertion | O(m) | O(m) |
| Word Search | O(m) | O(1) |
| Prefix Search | O(m) | O(1) |
| Deletion | O(m) | O(1) |
| Storage | — | O(ALPHABET_SIZE * N * M) |

*m — word length.
**N — number of words, M — average word length, ALPHABET_SIZE — alphabet size.
*/

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie - prefix tree structure
type Trie struct {
	root *TrieNode
}

// NewTrie - constructor
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

// Insert - inserts a word
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// Search - searches for a word
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

// StartsWith - verifies prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return true
}
