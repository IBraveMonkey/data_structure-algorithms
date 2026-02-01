package trie

/*
Trie (Префиксное дерево)
Trie (произносится как "try") — это дерево, в котором каждый узел хранит символ, а путь от корня до узла формирует префикс строки. Используется для эффективного хранения и поиска строк, особенно с общими префиксами.

Смысл и применение

Смысл: Быстрый поиск строк по префиксу или полному совпадению, экономия памяти за счет хранения общих префиксов.

Когда использовать:
Автодополнение (например, в поисковых системах).
Проверка орфографии.
Хранение словаря.
Алгоритмы сжатия (например, в IP-маршрутизации).
Поиск подстрок или префиксов.

Вставка: O(m), где m — длина слова.
Поиск слова: O(m).
Поиск префикса: O(m).
Память: O(ALPHABET_SIZE * N * M), где N — количество слов, M — средняя длина слова, ALPHABET_SIZE — размер алфавита.
*/

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie - структура префиксного дерева
type Trie struct {
	root *TrieNode
}

// NewTrie - конструктор
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

// Insert - вставка слова
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

// Search - поиск слова
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

// StartsWith - проверка префикса
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
