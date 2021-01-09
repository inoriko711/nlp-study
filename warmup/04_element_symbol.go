package warmup

import (
	"fmt"
	"strings"
)

// ElementSymbol は与えられた文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，それ以外の単語は先頭の1文字目と2文字目を取り出し，
// 取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成します
func ElementSymbol(sentence string) {
	words := strings.Fields(sentence)

	// TODO mapでは格納順を保持できないため、slice に変更する
	respMap := map[string]string{}

	for _, word := range words {
		strList := []rune(word)
		// TODO 先頭の1文字か、先頭の2文字かを判断する処理を入れる
		respMap[word] = string(strList[0])
	}

	fmt.Print(respMap)
}
