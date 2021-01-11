package warmup

func Aggregation(sentence1, sentence2 string) {
	strList1 := []rune(sentence1)
	strList2 := []rune(sentence2)
	ngram1 := ngram(strList1)
	ngram2 := ngram(strList2)

	// TODO 和集合を求める
	// TODO 積集合を求める
	// TODO 差集合を求める
	// TODO seが含まれてるかどうか調べる
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
