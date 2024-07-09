package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

//

func WordCount(s string) map[string]int {
	// Over here we import strings package
	//func which helps us separate words(one or more whitespaces)
	wordCounts := make(map[string]int)
	words := strings.Fields("string")
	for i := 0; i < len(words); i++ {
		wordCounts[words[i]] = wordCounts[words[i]] + 1
	}

	return wordCounts
}

func main() {
	wc.Test(WordCount)
}
