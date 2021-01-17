package warmup

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Typoglycemia(sentence string) {
	rand.Seed(time.Now().UnixNano())

	// 単語のリスト
	words := strings.Fields(sentence)

	// runeに直した単語のリスト
	runeList := make([][]rune, len(words))
	for i, word := range words {
		runeList[i] = []rune(word)
	}

	// 文字の順序をランダムに並び替えるプログラム
	for _, r := range runeList {
		if len(r) > 4 {
			// ss : シャッフルされる想定のrune
			ss := r[1 : len(r)-2]
			rand.Shuffle(len(ss), func(i, j int) { ss[i], ss[j] = ss[j], ss[i] })
			for k, s := range ss {
				r[k+1] = s
			}
		}
	}

	// runeを文字列に直したリスト
	resWordList := make([]string, len(words))
	for i, r := range runeList {
		resWordList[i] = string(r)
	}

	// 単語を文に戻す
	var resSentence string
	for _, word := range resWordList {
		resSentence += word + " "
	}
	fmt.Println("1.09. " + resSentence)
}
