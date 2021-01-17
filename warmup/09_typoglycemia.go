package warmup

import (
	"fmt"
	"strings"
)

func Typoglycemia(sentence string) {
	// 単語のリスト
	words := strings.Fields(sentence)

	// runeに直した単語のリスト
	runeList := make([][]rune, len(words))
	for i, word := range words {
		runeList[i] = []rune(word)
	}

	// TODO 文字の順序をランダムに並び替えるプログラム

	// runeを文字列に直したリスト
	resWordList := make([]string, len(words))
	for _, rune := range runeList {
		resWordList = append(resWordList, string(rune))
	}

	// 単語を文に戻す
	var resSentence string
	for _, word := range resWordList {
		resSentence += word + " "
	}
	fmt.Println("1.09. " + resSentence)
}
