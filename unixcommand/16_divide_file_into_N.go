package unixcommand

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

// https://nlp100.github.io/ja/ch02.html#16-%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%92n%E5%88%86%E5%89%B2%E3%81%99%E3%82%8B
func DivideFileIntoN(n int) {
	result := checkDivideFileIntoN(n)

	fmt.Printf("2.16. %t\n", result)
}

func checkDivideFileIntoN(n int) bool {
	//　コマンド実行
	out, err := exec.Command("split", "-l", fmt.Sprint(n), filepath.Join(Folder, FileName), filepath.Join(Folder, fmt.Sprintf("%s-split.", FileName))).Output()
	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println(out)
	return false
}
