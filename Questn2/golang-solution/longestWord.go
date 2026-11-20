package main

import (
	"fmt"
	"strings"
)

func longestWord(sentence string) string {
	words := strings.Fields(sentence)
	wordMapLength := make(map[string]int)
	maxLen := len(words[0])
	maxStr := words[0]
	for _, word := range words {
		wordMapLength[word] = len(word)

		if len(word) >= maxLen {
			maxLen = len(word)
			maxStr = word
		}
	}
	fmt.Printf("%#v\n", wordMapLength)
	return maxStr
}

func main() {
	sentence := "I am a very linguistic person"
	fmt.Println(longestWord(sentence))
	sentence2 := "have a nice day"
	fmt.Println(longestWord(sentence2))
	sentence3 := "the quick brown fox jumped over the lazy dog"
	fmt.Println(longestWord(sentence3))
	sentence4 := "who did eat the ham"
	fmt.Println(longestWord(sentence4))
	sentence5 := "potato"
	fmt.Println(longestWord(sentence5))
}
