package unixcommand

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
)

// https://nlp100.github.io/ja/ch02.html#16-%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%92n%E5%88%86%E5%89%B2%E3%81%99%E3%82%8B
func DivideFileIntoN(n int) {
	if n < 50 {
		log.Println("2.16. 50以上の整数を指定してください")
		return
	}

	// ファイルの内容を取得する
	fileList, err := getFileElements(filepath.Join(Folder, FileName))
	if err != nil {
		log.Println(err)
		return
	}

	// ファイルをN行ずつに分けて保持する
	var nLineFile []string
	var nLineFiles [][]string
	for i := 0; i < len(fileList); i = i + n {
		nLineFile = fileList[i : i+n]
		nLineFiles = append(nLineFiles, nLineFile)
	}

	result := checkDivideFileIntoN(n, nLineFiles)
	fmt.Printf("2.16. %t\n", result)
}

func checkDivideFileIntoN(n int, gots [][]string) bool {
	// コマンド実行結果を一時的に保存するフォルダを準備する
	tmpFolder := Folder + "/tmp"
	if err := os.MkdirAll(tmpFolder, 0777); err != nil {
		log.Println(err)
		return false
	}
	//　コマンド実行
	_, err := exec.Command("split", "-l", fmt.Sprint(n), filepath.Join(Folder, FileName), filepath.Join(tmpFolder, fmt.Sprintf("%s-split.", FileName))).Output()
	if err != nil {
		log.Println(err)
		return false
	}

	// splitで生成したファイル一覧を取得する
	files, err := ioutil.ReadDir(tmpFolder)
	if err != nil {
		panic(err)
	}

	// 結果の確認
	for i, got := range gots {
		// ファイルを開いて中身を取得する
		want, err := getFileElements(filepath.Join(tmpFolder, files[i].Name()))
		if err != nil {
			log.Println(err)
			return false
		}
		want = revoke(want, "")
		got = revoke(got, "")

		if !reflect.DeepEqual(want, got) {
			return false
		}
	}

	//TODO 生成されたファイルを削除する仕組みを作る
	return true
}
