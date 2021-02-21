package unixcommand

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

// DifferentString 1列目の文字列の種類（異なる文字列の集合）を求める関数
func DifferentString() bool {

	// ファイルの中身を読み込む
	f, err := os.Open(filepath.Join(UnixcommandFolder, PopularNamesFileName))
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer f.Close()

	// 1列目のみ取得する
	col1Words := make(map[string]bool)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, "\t")
		if !col1Words[words[0]] {
			col1Words[words[0]] = true
		}
	}

	var result []string
	for word := range col1Words {
		result = append(result, word)
	}
	sort.Strings(result)

	err = checkDifferentString(result)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// 結果の出力
	// fmt.Println(result)

	fmt.Println("2.17. ok")

	return true
}

func checkDifferentString(got []string) error {
	// 実際のコマンドを準備'cut -f 1 unixcommand/popular-names.txt | sort | uniq'
	cmdstr := fmt.Sprintf("cut -f 1 %s | sort | uniq", filepath.Join(UnixcommandFolder, PopularNamesFileName))

	// 実行結果の取得
	out, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		return err
	}
	want := strings.Split(string(out), "\n")

	sort.Strings(want)
	want = revoke(want, "")

	// 結果が相違ないことの確認
	if !reflect.DeepEqual(want, got) {
		return fmt.Errorf(fmt.Sprintf("want %s\n got %s", want, got))
	}

	return nil
}
