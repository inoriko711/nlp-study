package unixcommand

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func OutputTheFirstNLines(n int) {
	// TODO 実装

	result := checkOutputTheFirstNLines(n)

	fmt.Printf("2.14. %t\n", result)
}

// TODO 動作確認
func checkOutputTheFirstNLines(n int) bool {
	//　コマンド実行
	out, err := exec.Command("head", fmt.Sprintf("-%d", n), filepath.Join("unixcommand", "popular-names.txt")).Output()
	if err != nil {
		log.Println(err)
		return false
	}
	outList := strings.Split(string(out), "\n")
	fmt.Println(outList)

	return false
}
