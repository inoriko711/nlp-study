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

	boolCol1 := checkResult("-1", "col1.txt")
	boolCol2 := checkResult("2-2", "col2.txt")

	fmt.Printf("2.12. col1.txt is %t, col2.txt is %t\n", boolCol1, boolCol2)
}

// 出力結果とコマンド実施結果が一致していることの確認
func checkResult(col, filename string) bool {
	out, err := exec.Command("cut", "-f", col, filepath.Join("unixcommand", "popular-names.txt")).Output()
	if err != nil {
		log.Fatal(err)
		return false
	}
	outList := strings.Split(string(out), "\n")

	f, err := os.Open(filepath.Join("unixcommand", filename))
	if err != nil {
		log.Fatal(err)
		return false
	}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		if scanner.Text() != string(outList[i]) {
			log.Fatalf("want: %s, got: %s", string(outList[i]), scanner.Text())
			return false
		}
		i++
	}
	return true
}

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
