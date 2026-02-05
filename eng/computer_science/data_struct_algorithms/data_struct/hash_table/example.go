package hash_table

import (
	"fmt"
	"sort"
)

// Example demonstrates the use of a hash table with various examples
func Example() {
	// Example 1: Find the extra letter
	a := "abcd"
	b := "abcde"
	result := ExtraLetter(a, b)
	fmt.Printf("Extra letter in '%s' compared to '%s': %s\n", b, a, result)

	// Example 2: Sum of two elements
	nums := []int{2, 7, 11, 15}
	target := 9
	indices := TwoSum(nums, target)
	fmt.Printf("Indices of two numbers for sum %d in %v: %v\n", target, nums, indices)

	// Example 3: Grouping anagrams
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagramGroups := GroupAnagrams(strs)
	fmt.Printf("Anagram groups for %v: %v\n", strs, anagramGroups)

	// Example 4: Counting frequency
	numsFreq := []int{1, 2, 3, 2, 1, 1}
	freqMap := CountFrequency(numsFreq)
	fmt.Printf("Number frequency for %v: %v\n", numsFreq, freqMap)

	// Example 5: Duplicate check
	hasDup := ContainsDuplicate([]int{1, 2, 3, 1})
	fmt.Printf("Duplicates in [1,2,3,1]: %t\n", hasDup)

	hasDup = ContainsDuplicate([]int{1, 2, 3, 4})
	fmt.Printf("Duplicates in [1,2,3,4]: %t\n", hasDup)
}

// Problem 1: Find the difference between two strings
// Two strings a and b are given.
// String b is formed by shuffling string a and adding one letter.
// Return that letter.
//
// Example: a = "abc", b = "beca" => 'e'
// Time complexity: O(n), space complexity: O(n)
func ExtraLetter(a, b string) string {
	hTable := make(map[rune]int, len(b))

	// Count the frequency of characters in string b
	for _, v := range b {
		hTable[v]++
	}

	// Decrement the counts for characters from string a
	for _, v := range a {
		hTable[v]--
	}

	// Find the character with a non-zero count
	for key, val := range hTable {
		if val > 0 {
			return string(key)
		}
	}

	return ""
}

// Problem 2: Sum of two array elements
// Given an unsorted array of integers and a target number.
// Write a function that finds two such elements in the array whose sum is equal to target.
//
// Example: data = [2, 7, 11, 15], target = 9 => [0, 1] (data[0] + data[1] = 2 + 7 = 9)
// Time complexity: O(n), space complexity: O(n)
func TwoSum(data []int, target int) []int {
	// Create a hash table to store values and their indices
	cache := make(map[int]int)

	for i, num := range data {
		complement := target - num // The complement needed to reach the target

		// Check if the complement is in the hash table
		if index, ok := cache[complement]; ok {
			return []int{index, i} // Return indices of the two elements
		}

		// Save the current element and its index
		cache[num] = i
	}

	return []int{} // If no pair is found
}

// Problem 3: Anagram Grouping
// Find and group anagram words together.
// An anagram is a word or phrase formed by rearranging the letters of another word or phrase.
//
// Example: ["eat", "tea", "tan", "ate", "nat", "bat"] => [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
// Time complexity: O(n * m * log(m)), where n is the number of strings and m is the average string length
func GroupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, str := range strs {
		// Sort the string's characters to get a common key for anagrams
		sortedStr := sortString(str)

		// Add the string to the group with the corresponding key
		anagramGroups[sortedStr] = append(anagramGroups[sortedStr], str)
	}

	// Convert the map to a slice of slices
	result := make([][]string, 0, len(anagramGroups))
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

// Helper function to sort characters in a string
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Problem 4: Element Frequency Count
// Count how many times each element occurs in an array
//
// Example: [1, 2, 3, 2, 1, 1] => {1: 3, 2: 2, 3: 1}
// Time complexity: O(n), space complexity: O(k), where k is the number of unique elements
func CountFrequency(nums []int) map[int]int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++ // Increment counter for the current element
	}

	return freqMap
}

// Problem 5: Duplicate Check
// Check if there are duplicate elements in an array
//
// Example: [1, 2, 3, 1] => true, [1, 2, 3, 4] => false
// Time complexity: O(n), space complexity: O(n)
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true // Duplicate found
		}
		seen[num] = true
	}

	return false // No duplicates
}

// Problem 6: Intersection of Two Arrays
// Find elements that are present in both arrays
//
// Example: [1, 2, 2, 1], [2, 2] => [2]
// Time complexity: O(n + m), space complexity: O(min(n, m))
func Intersection(nums1, nums2 []int) []int {
	set1 := make(map[int]bool)
	for _, num := range nums1 {
		set1[num] = true
	}

	set2 := make(map[int]bool)
	for _, num := range nums2 {
		if set1[num] {
			set2[num] = true
		}
	}

	result := make([]int, 0, len(set2))
	for num := range set2 {
		result = append(result, num)
	}

	return result
}

// Problem 7: First Non-Repeating Character
// Find the first character that appears only once in a string
//
// Example: "leetcode" => 'l', "loveleetcode" => 'v'
// Time complexity: O(n), space complexity: O(k), where k is the number of unique characters
func FirstUniqueChar(s string) int {
	charCount := make(map[rune]int)

	// Count the frequency of each character
	for _, char := range s {
		charCount[char]++
	}

	// Find the first character with a frequency of 1
	for i, char := range s {
		if charCount[char] == 1 {
			return i
		}
	}

	return -1 // No unique character
}

// Problem 8: Palindrome Permutation check
// Check if string characters can be rearranged to form a palindrome
//
// Example: "aab" => true ("aba"), "carerac" => true ("racecar")
// Time complexity: O(n), space complexity: O(k), where k is the number of unique characters
func CanPermutePalindrome(s string) bool {
	charCount := make(map[rune]int)

	// Count the frequency of each character
	for _, char := range s {
		charCount[char]++
	}

	// Count the number of characters with an odd frequency
	oddCount := 0
	for _, count := range charCount {
		if count%2 == 1 {
			oddCount++
		}
	}

	// A palindrome is possible if no more than one character has an odd frequency
	return oddCount <= 1
}

// Problem 9: Contains Duplicate II
// Given an integer array nums and an integer k. Return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
func ContainsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)

	for i, num := range nums {
		if prevIndex, exists := indexMap[num]; exists {
			if i-prevIndex <= k {
				return true
			}
		}
		// Update the last seen index for this number
		indexMap[num] = i
	}

	return false
}

// Problem 10: Valid Sudoku
// Determine if a 9x9 Sudoku board is valid. Only filled cells need to be validated.
func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	// Initialize maps
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := board[i][j]
			if cell == '.' {
				continue
			}

			// Row check
			if rows[i][cell] {
				return false
			}
			rows[i][cell] = true

			// Column check
			if cols[j][cell] {
				return false
			}
			cols[j][cell] = true

			// 3x3 square check
			boxIndex := (i/3)*3 + j/3
			if boxes[boxIndex][cell] {
				return false
			}
			boxes[boxIndex][cell] = true
		}
	}

	return true
}

// Problem 11: Group Shifted Strings
// We can shift a string right to get a new string.
// For example, "abc" -> "bcd". A string is grouped with others that can be shifted to form each other.
func GroupShiftedStrings(strings []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strings {
		// Create a key based on relative shifts from the first character
		key := ""
		if len(s) > 0 {
			base := rune(s[0])

			for _, c := range s {
				// Calculate the relative shift, handling the 26-character cycle
				shift := (c - base + 26) % 26
				key += string(shift) + ","
			}
		}

		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
