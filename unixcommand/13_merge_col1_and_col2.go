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

func MergeCol1AndCol2() {
	// ファイルの内容を取得する
	col1, err := os.Open(filepath.Join("unixcommand", "col1.txt"))
	if err != nil {
		log.Println(err)
		return
	}
	scanner1 := bufio.NewScanner(col1)

	col2, err := os.Open(filepath.Join("unixcommand", "col2.txt"))
	if err != nil {
		log.Println(err)
		return
	}
	scanner2 := bufio.NewScanner(col2)

	var resList []string
	for scanner1.Scan() {
		scanner2.Scan()
		resList = append(resList, fmt.Sprintf("%s\t%s", scanner1.Text(), scanner2.Text()))
	}

	res := checkMergeResult(resList)
	fmt.Printf("2.13. %t\n", res)
}

func checkMergeResult(got []string) bool {
	//　コマンド実行
	out, err := exec.Command("paste", filepath.Join("unixcommand", "col1.txt"), filepath.Join("unixcommand", "col2.txt")).Output()
	if err != nil {
		log.Println(err)
		return false
	}
	outList := strings.Split(string(out), "\n")

	// if len(outList) != len(got) {
	// 	log.Printf("want len %d, got len %d\n", len(outList), len(got))
	// 	return false
	// }

	for i, str := range got {
		if str != outList[i] {
			log.Printf("want %s, got %s\n", outList[i], str)
			return false
		}
	}

	return true
}
