package main

import (
	"fmt"
)

func KUniqueCharacters(str string, k int) string {
	tempWord := str[0 : k+1]
	word := tempWord
	wordEnd := k

	for wordEnd < len(str)-1 {
		if len(uniqueChars(tempWord)) <= k {
			if len(tempWord) > len(word) {
				word = tempWord
			}
			wordEnd++
			tempWord += string(str[wordEnd])
		} else {
			tempWord = tempWord[1:]
		}
	}

	if len(tempWord) > len(word) {
		word = tempWord
	}

	return word
}

func uniqueChars(str string) []byte {
	var unique []byte
	seen := make(map[byte]bool)

	for i := 0; i < len(str); i++ {
		char := str[i]
		if !seen[char] {
			seen[char] = true
			unique = append(unique, char)
		}
	}

	return unique
}

func main() {
	str := "abcd"
	k := 3

	result := KUniqueCharacters(str, k)
	fmt.Println(result) // Output: bbcdd
}
