package warmup

import "fmt"

func Aggregation(sentence1, sentence2 string) {
	strList1 := []rune(sentence1)
	strList2 := []rune(sentence2)
	ngram1 := ngram(strList1)
	ngram2 := ngram(strList2)

	fmt.Print("1.06. 和集合: ")
	fmt.Println(union(ngram1, ngram2))

	// TODO 積集合を求める
	fmt.Print("1.06. 積集合: ")
	// fmt.Println(union(ngram1, ngram2))

	// TODO 差集合を求める
	fmt.Print("1.06. 差集合: ")
	// fmt.Println(union(ngram1, ngram2))

	// TODO seが含まれてるかどうか調べるs
}
func union(ss ...[]string) []string {
	m := map[string]int{}
	for _, s := range ss {
		for _, v := range s {
			m[v]++ // 出現回数のカウント
		}
	}
	res := []string{}
	for str, c := range m {
		if c == 1 {
			// 出現回数が１回のものだけを抽出
			res = append(res, str)
		}
	}
	return res
}

func ngram(runeList []rune) []string {
	var respList []string
	var str string
	for i, r := range runeList {
		str = string(r)
		if i < len(runeList)-1 {
			i++
			str = str + string(runeList[i])
		}
		respList = append(respList, str)
	}
	return respList
}
