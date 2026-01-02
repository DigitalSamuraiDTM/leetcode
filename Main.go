package main

import (
	"strings"
)

func main() {
	//println(wordPattern("abba", "dog cat cat dog"))
	//println(wordPattern("abba", "dog cat cat fish"))
	//println(wordPattern("aaaa", "dog cat cat dog"))
	//println(wordPattern("abba", "dog dog dog dog"))
	println(wordPattern("aaa", "aa aa aa aa"))
}

func wordPattern(pattern string, s string) bool {
	patternCharToWord := make(map[uint8]string)
	patternWordToChar := make(map[string]uint8)
	splitedWords := strings.Split(s, " ")
	if len(splitedWords) != len(pattern) {
		return false
	}
	for index, word := range splitedWords {
		char := pattern[index]
		patternedWord, isWordExist := patternCharToWord[char]
		patternedChar, isCharExist := patternWordToChar[word]
		if !isWordExist && !isCharExist {
			patternCharToWord[char] = word
			patternWordToChar[word] = char
			continue
		}
		if patternedWord != word || patternedChar != char {
			return false
		}
	}
	return true
}
