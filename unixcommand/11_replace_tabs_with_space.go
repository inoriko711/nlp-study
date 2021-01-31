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

func ReplaceTabsWithSpace() {
	// ファイル読み込み
	f, err := os.Open(filepath.Join(Folder, FileName))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	// 置換後文字列の取得
	replacedStr := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "\t", " ")
		replacedStr = append(replacedStr, line)
	}

	// 答え合わせ
	// コマンドの実行
	out, err := exec.Command("sed", "-E", "s/\t/ /g", filepath.Join(Folder, FileName)).Output()
	if err != nil {
		log.Fatal(err)
		return
	}
	outList := strings.Split(string(out), "\n")

	var result []string
	for i, str := range replacedStr {
		if outList[i] != str {
			result = append(result, fmt.Sprintf("want %s, got %s\n", outList[i], str))
		}
	}

	if len(result) == 0 {
		fmt.Println("2.11. ok")
		return
	}

	for _, res := range result {
		fmt.Printf("2.11. %s", res)
	}

}
