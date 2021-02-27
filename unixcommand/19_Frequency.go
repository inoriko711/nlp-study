package unixcommand

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// 各行の1列目の文字列の出現頻度を求め，その高い順に並べて表示せよ．
// 確認にはcut, uniq, sortコマンドを用いよ．
func Frequency() bool {
	// ファイルを開く
	f, err := os.Open(filepath.Join(UnixcommandFolder, PopularNamesFileName))
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer f.Close()

	// 名をキー、出現頻度を値に持つMap
	nameFrequencyMap := make(map[string]int)

	// ファイルの中身を読み取って、nameFrequencyMapに値をつめる
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, "\t")

		// 1列目を取得
		name := words[0]

		nameFrequencyMap[name] += 1
	}

	// ソートするために配列に書き出す
	resList := make([]string, len(nameFrequencyMap))
	for key, val := range nameFrequencyMap {
		resList = append(resList, fmt.Sprintf("%d\t%s", val, key))
	}

	sort.Slice(resList, func(i, j int) bool {
		nameFrequencyListI := strings.Split(resList[i], "\t")
		frequencyI, _ := strconv.Atoi(nameFrequencyListI[0])
		nameFrequencyListJ := strings.Split(resList[j], "\t")
		frequencyJ, _ := strconv.Atoi(nameFrequencyListJ[0])
		return frequencyI > frequencyJ
	})

	err = checkFrequency(resList)
	if err != nil {
		log.Println(err)
		return false
	}

	// 結果の出力
	// fmt.Println(resList)

	fmt.Println("2.19 ok")
	return false
}

func checkFrequency(got []string) error {
	// cut -f 1 popular-names.txt | sort | uniq -c | sort -nr
	cmdstr := fmt.Sprintf("cut -f 1 %s | sort | uniq -c | sed -e 's/ *\\([0-9]*\\) /\\1\t/' | sort -nr", filepath.Join(UnixcommandFolder, PopularNamesFileName))

	// 実行結果の取得
	out, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		return err
	}

	want := strings.Split(string(out), "\n")
	want = revoke(want, "")

	fmt.Println(want)
	fmt.Println(got)

	// TODO 実装する
	return errors.New("未実装")
}
