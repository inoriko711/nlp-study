package warmup

import (
	"fmt"
	"strings"
)

// ElementSymbol は与えられた文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，それ以外の単語は先頭の1文字目と2文字目を取り出し，
// 取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成します
func ElementSymbol(sentence string) {
	// 先頭の1文字のみとる順番を保持するリスト
	firstStrList := []int{0, 4, 5, 6, 7, 8, 14, 15, 18}

	words := strings.Fields(sentence)

	respMap := map[string]string{}

	j := 0
	for i, word := range words {
		strList := []rune(word)

		// 先頭一文字のみか、二文字かを判断して格納する
		if i == firstStrList[j] {
			respMap[word] = string(strList[0])
			if j < len(firstStrList)-1 {
				j++
			}
		} else {
			respMap[word] = string(strList[0]) + string(strList[1])
		}
	}
	fmt.Print("1.04. ")
	for _, word := range words {
		fmt.Print(word + ":" + respMap[word] + " ")
	}
}
