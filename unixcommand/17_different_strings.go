package unixcommand

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// DifferentString 1列目の文字列の種類（異なる文字列の集合）を求める関数
func DifferentString(file string) []string {

	// ファイルの中身を読み込む
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// 1列目のみ取得する
	col1Words := make(map[string]bool)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, "\t")
		if !col1Words[words[0]] {
			col1Words[words[0]] = true
		}
	}

	var result []string
	for word := range col1Words {
		result = append(result, word)
	}
	sort.Strings(result)
	fmt.Printf("2.17. 1列目の文字列の種類： %s\n", result)

	return result
}
