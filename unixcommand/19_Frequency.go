package unixcommand

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

// 各行の1列目の文字列の出現頻度を求め，その高い順に並べて表示せよ．
// 確認にはcut, uniq, sortコマンドを用いよ．
func Frequency() bool {
	// TODO 実装する

	err := checkFrequency()
	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println("2.19 ok")
	return false
}

func checkFrequency() error {
	// cut -f 1 popular-names.txt | sort | uniq -c | sort -nr
	cmdstr := fmt.Sprintf("cut -f 1 %s | sort | uniq -c | sort -nr", filepath.Join(UnixcommandFolder, PopularNamesFileName))

	// 実行結果の取得
	out, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		return err
	}

	want := strings.Split(string(out), "\n")

	fmt.Println(want)

	// TODO 実装する
	return errors.New("未実装")
}
