package warmup

import "fmt"

//与えられた文字列の各文字を，以下の仕様で変換する関数
// 英小文字ならば(219 - 文字コード)の文字に置換
// その他の文字はそのまま出力
func Cipher(str string) {
	strList := []rune(str)

	cipher := func(list []rune) string {
		for i, s := range strList {
			if 'a' < s && s < 'z' {
				strList[i] = 219 - s
			}
		}
		return string(list)
	}

	fmt.Println("1.08. 暗号化: " + cipher(strList))
	fmt.Println("1.08. 複合化: " + cipher(strList))
}
