package trie

import "fmt"

// Example demonstrates the use of a Trie
func Example() {
	// Create a new prefix tree
	trie := NewTrie()

	// Insert words
	words := []string{"apple", "app", "application", "banana"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Check for word existence
	tests := []string{"apple", "app", "appl", "banana", "orange"}
	for _, word := range tests {
		exists := trie.Search(word)
		fmt.Printf("Does word '%s' exist? %t\n", word, exists)
	}

	// Check prefixes
	prefixes := []string{"app", "ban", "ora"}
	for _, prefix := range prefixes {
		startsWith := trie.StartsWith(prefix)
		fmt.Printf("Are there words with prefix '%s'? %t\n", prefix, startsWith)
	}
}

// Task: Longest Word in Dictionary
// Given an array of strings words representing a dictionary.
// Find the longest word in dictionary that can be built one character at a time,
// where each next part must also be a word in the dictionary.
//
// Example: words = ["w","wo","wor","worl","world"] => "world"
// Note: To solve this problem, you can use a Trie where each node stores whether it's an end of some word.
func LongestWord(words []string) string {
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}

	longest := ""

	var dfs func(node *TrieNode, currentWord string)
	dfs = func(node *TrieNode, currentWord string) {
		if len(currentWord) > len(longest) || (len(currentWord) == len(longest) && currentWord < longest) {
			longest = currentWord
		}

		for char, child := range node.children {
			// Continue only if current node is the end of a word (i.e., word built one letter at a time)
			if child.isEnd {
				dfs(child, currentWord+string(char))
			}
		}
	}

	// Run DFS from the root. But in current Trie implementation root is not the end of a word.
	// We need to modify the approach as root is empty.
	// Let's traverse the root's children that are words.

	for char, child := range trie.root.children {
		if child.isEnd {
			dfs(child, string(char))
		}
	}

	return longest
}
