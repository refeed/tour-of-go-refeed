package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCountMap := make(map[string]int)
	splittedWords := strings.Fields(s)

	for _, word := range splittedWords {
		count, exists := wordCountMap[word]
		if !exists {
			wordCountMap[word] = 1
			continue
		}
		wordCountMap[word] = count + 1
	}

	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
