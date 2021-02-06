package unixcommand

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// https://nlp100.github.io/ja/ch02.html#16-%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%92n%E5%88%86%E5%89%B2%E3%81%99%E3%82%8B
func DivideFileIntoN(n int) {
	if n < 50 {
		fmt.Println("2.16. 50以上の整数を指定してください")
		return
	}

	// ファイルの内容を取得する
	f, err := os.Open(filepath.Join(Folder, FileName))
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	// ファイルの中身を全て取得する
	scanner := bufio.NewScanner(f)
	var fileList []string
	for scanner.Scan() {
		fileList = append(fileList, scanner.Text())
	}

	// ファイルをN行ずつに分けて保持する
	var nLineFile []string
	var nLineFiles [][]string
	for i := 0; i < len(fileList); i = i + n {
		nLineFile = fileList[i : i+n]
		nLineFiles = append(nLineFiles, nLineFile)
	}

	result := checkDivideFileIntoN(n, nLineFiles)
	fmt.Printf("2.16. %t\n", result)
}

func checkDivideFileIntoN(n int, got [][]string) bool {
	//　コマンド実行
	out, err := exec.Command("split", "-l", fmt.Sprint(n), filepath.Join(Folder, FileName), filepath.Join(Folder, fmt.Sprintf("%s-split.", FileName))).Output()
	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println(out)

	for _, g := range got {
		fmt.Println(len(g))
	}

	//TODO 生成されたファイルを削除する仕組みを作る
	return false
}
