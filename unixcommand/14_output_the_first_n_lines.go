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

func OutputTheFirstNLines(n int) {
	// ファイルの内容を取得する
	f, err := os.Open(filepath.Join(UnixcommandFolder, PopularNamesFileName))
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	// ファイルの先頭N行を取得する
	scanner := bufio.NewScanner(f)
	var resList []string
	for i := 0; scanner.Scan() && i < n; i++ {
		resList = append(resList, scanner.Text())
	}

	result := checkOutputTheFirstNLines(n, resList)

	fmt.Printf("2.14. %t\n", result)
}

func checkOutputTheFirstNLines(n int, got []string) bool {
	//　コマンド実行
	out, err := exec.Command("head", fmt.Sprintf("-%d", n), filepath.Join(UnixcommandFolder, PopularNamesFileName)).Output()
	if err != nil {
		log.Println(err)
		return false
	}
	want := strings.Split(string(out), "\n")
	want = revoke(want, "")

	return reflect.DeepEqual(want, got)
}
