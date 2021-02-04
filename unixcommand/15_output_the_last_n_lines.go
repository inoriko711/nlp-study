package unixcommand

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

func OutputTheLastNLines(n int) {
	// ファイルの内容を取得する
	f, err := os.Open(filepath.Join(Folder, FileName))
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	// ファイルの中身を全て取得する
	scanner := bufio.NewScanner(f)
	var resList []string
	for scanner.Scan() {
		resList = append(resList, scanner.Text())
	}

	// ファイルの最後のN行を取得する
	resList = resList[len(resList)-n:]

	result := checkOutputTheLastNLines(n, resList)
	fmt.Printf("2.15. %t\n", result)
}
func checkOutputTheLastNLines(n int, got []string) bool {
	//　コマンド実行
	out, err := exec.Command("tail", fmt.Sprintf("-%d", n), filepath.Join(Folder, FileName)).Output()
	if err != nil {
		log.Println(err)
		return false
	}

	want := strings.Split(string(out), "\n")
	want = revoke(want, "")

	return reflect.DeepEqual(want, got)
}
