package unixcommand

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReplaceTabsWithSpaces() {
	f, err := os.Open(filepath.Join("unixcommand", "popular-names.txt"))
	if err != nil {
		log.Fatal(err)
		return
	}

	// 1列目と2列目のみ取得する
	var col1Words []string
	var col2Words []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, "\t")
		col1Words = append(col1Words, words[0])
		col2Words = append(col2Words, words[1])
	}

	// 各行の1列目だけを抜き出したファイルを作成する
	// % cut -f -1 unixcommand/popular-names.txt
	col1File, err := os.OpenFile(filepath.Join("unixcommand", "col1.txt"), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer col1File.Close()
	for _, word := range col1Words {
		_, err := col1File.WriteString(word + "\n")
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	// 各行の2列目だけを抜き出したファイルを作成する
	// % cut -f 2-2 unixcommand/popular-names.txt
	col2File, err := os.OpenFile(filepath.Join("unixcommand", "col2.txt"), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer col2File.Close()
	for _, word := range col2Words {
		_, err := col2File.WriteString(word + "\n")
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	fmt.Println("2.12. check col1.txt and col2.txt")
}
