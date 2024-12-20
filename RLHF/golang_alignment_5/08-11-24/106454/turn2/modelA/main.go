package main

import (
	"fmt"
	"strings"
)

// isVowel checks if a character is a vowel
func isVowel(c rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, c)
}

// countVowelsInString counts the vowels in a given string
func countVowelsInString(input string) int {
	return len(strings.Filter(input, isVowel))
}

// main function
func main() {
	testStrings := []string{"Hello", "world", "Go is great", "Functional programming rocks!"}

	// Using map to apply the countVowelsInString function to each test string
	vowelCounts := map(testStrings, countVowelsInString)

	for _, str := range testStrings {
		fmt.Printf("Vowels in '%s': %d\n", str, vowelCounts[str])
	}
}

// map function implementation
func map(strings []string, f func(string) int) map[string]int {
	result := make(map[string]int)
	for _, str := range strings {
		result[str] = f(str)
	}
	return result
}
