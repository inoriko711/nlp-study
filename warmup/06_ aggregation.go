package warmup

import (
	"fmt"
)

func Aggregation(sentence1, sentence2 string) {
	strList1 := []rune(sentence1)
	strList2 := []rune(sentence2)
	ngram1 := ngram(strList1)
	ngram2 := ngram(strList2)

	fmt.Print("1.06. ngram: ")
	fmt.Println(ngram1)
	fmt.Print("1.06. ngram: ")
	fmt.Println(ngram2)

	fmt.Print("1.06. 和集合: ")
	fmt.Println(union(ngram1, ngram2))

	fmt.Print("1.06. 積集合: ")
	fmt.Println(intersection(ngram1, ngram2))

	fmt.Print("1.06. 差集合: ")
	fmt.Println(differenceSet(ngram1, ngram2))

	fmt.Printf("1.06. seの有無: %s→%t, %s→%t\n", sentence1, exist(ngram1, "se"), sentence2, exist(ngram2, "se"))
}

// 差集合
func differenceSet(ss ...[]string) []string {
	m := map[string]int{}
	for _, s := range ss {
		// 重複削除
		s = union(s)
		for _, v := range s {
			m[v]++
		}
	}

	var res []string
	for v, i := range m {
		if i == 1 {
			res = append(res, v)
		}
	}
	return res
}

// 積集合
func intersection(ss ...[]string) []string {
	m := map[string]int{}
	for _, s := range ss {
		// 重複削除
		s = union(s)
		for _, v := range s {
			m[v]++
		}
	}

	var res []string
	for v, i := range m {
		if i == len(ss) {
			res = append(res, v)
		}
	}
	return res
}

// 重複確認
func exist(strList []string, searchedWord string) bool {
	for _, str := range strList {
		if str == searchedWord {
			return true
		}
	}
	return false
}

// 和集合
func union(ss ...[]string) []string {
	var res []string
	for _, s := range ss {
		for _, v := range s {
			if !exist(res, v) {
				res = append(res, v)
			}
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
			respList = append(respList, str)
		}
	}
	return respList
}
