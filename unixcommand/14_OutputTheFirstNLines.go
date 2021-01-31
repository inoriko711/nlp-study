package unixcommand

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func OutputTheFirstNLines(n int) {
	// TODO 実装
	// ファイルの内容を取得する
	col1, err := os.Open(filepath.Join(Folder, "col1.txt"))
	if err != nil {
		log.Println(err)
		return
	}
	defer col1.Close()

	result := checkOutputTheFirstNLines(n)

	fmt.Printf("2.14. %t\n", result)
}

// TODO 動作確認
func checkOutputTheFirstNLines(n int) bool {
	//　コマンド実行
	out, err := exec.Command("head", fmt.Sprintf("-%d", n), filepath.Join(Folder, FileName)).Output()
	if err != nil {
		log.Println(err)
		return false
	}
	outList := strings.Split(string(out), "\n")
	fmt.Println(outList)

	return false
}
