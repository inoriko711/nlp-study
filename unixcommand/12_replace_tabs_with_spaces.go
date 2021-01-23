package unixcommand

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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
	createFiles("col1.txt", col1Words)

	// 各行の2列目だけを抜き出したファイルを作成する
	// % cut -f 2-2 unixcommand/popular-names.txt
	createFiles("col2.txt", col2Words)

	out, err := exec.Command("cut", "-f", "-1", "unixcommand/popular-names.txt").Output()
	if err != nil {
		log.Fatal(err)
		return
	}
	outList := strings.Split(string(out), "\n")

	f1, err := os.Open(filepath.Join("unixcommand", "col1.txt"))
	if err != nil {
		log.Fatal(err)
		return
	}
	scanner = bufio.NewScanner(f1)
	i := 0
	for scanner.Scan() {
		if scanner.Text() != string(outList[i]) {
			log.Fatalf("want: %s, got: %s", string(outList[i]), scanner.Text())
			break
		}
		i++
	}

	fmt.Println("2.12. check col1.txt and col2.txt")
}

// 出力結果とコマンド実施結果が一致していることの確認

// ファイル名と文字列を受け取り、一行ずつ保存するメソッド
func createFiles(fileName string, words []string) {
	colFile, err := os.OpenFile(filepath.Join("unixcommand", fileName), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer colFile.Close()
	for _, word := range words {
		_, err := colFile.WriteString(word + "\n")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
