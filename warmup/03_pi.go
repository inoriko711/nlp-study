package warmup

import (
	"fmt"
	"strings"
)

// Pi は文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを出力します。
func Pi(sentence string) {
	words := strings.Fields(sentence)
	var respList []int
	for _, word := range words {
		respList = append(respList, len(word)) // len(string)はバイト長を返すので、日本語に対しては使用できない
	}
	fmt.Printf("1.03. %d\n", respList)
}
