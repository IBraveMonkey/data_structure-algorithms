package trie

import "fmt"

// Example демонстрирует использование Trie
func Example() {
	// Создаем новое префиксное дерево
	trie := NewTrie()

	// Вставляем слова
	words := []string{"apple", "app", "application", "banana"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Проверяем наличие слов
	tests := []string{"apple", "app", "appl", "banana", "orange"}
	for _, word := range tests {
		exists := trie.Search(word)
		fmt.Printf("Слово '%s' существует? %t\n", word, exists)
	}

	// Проверяем префиксы
	prefixes := []string{"app", "ban", "ora"}
	for _, prefix := range prefixes {
		startsWith := trie.StartsWith(prefix)
		fmt.Printf("Есть слова с префиксом '%s'? %t\n", prefix, startsWith)
	}
}

// Задача: Самое длинное слово в словаре
// Дан массив строк words, представляющий словарь.
// Найти самое длинное слово в словаре, которое может быть построено по одной букве за раз,
// где каждая следующая часть также должна быть словом в словаре.
//
// Пример: words = ["w","wo","wor","worl","world"] => "world"
// Примечание: Для решения этой задачи можно использовать Trie, где каждый узел будет хранить информацию, является ли он концом какого-либо слова.
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
			// Продолжаем только если текущий узел является концом слова (т.е. слово построено по одной букве)
			if child.isEnd {
				dfs(child, currentWord+string(char))
			}
		}
	}

	// Запускаем DFS от корня. Но в текущей реализации Trie root не является концом слова.
	// Нам нужно модифицировать подход, так как root пустой.
	// Пройдемся по детям корня, которые являются словами.

	for char, child := range trie.root.children {
		if child.isEnd {
			dfs(child, string(char))
		}
	}

	return longest
}
