package warmup

import (
	"fmt"
	"strings"
)

func Ngram(sentence string) {

	wordList := strings.Fields(sentence)
	strList := []rune(sentence)

	fmt.Print("1.05. 単語bi-gram:")
	for i, word := range wordList {
		fmt.Print(word)
		if i+1 < len(wordList) {
			fmt.Print(wordList[i+1])
			fmt.Print(" ")
		}
	}

	fmt.Print("\n1.05. 文字bi-gram:")
	for i, str := range strList {
		fmt.Print(string(str))
		if i+1 < len(strList) {
			fmt.Print(string(strList[i+1]))
			fmt.Print(" ")
		}
	}

	fmt.Println()
}
