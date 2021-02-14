package unixcommand

import (
	"bufio"
	"os"
	"path/filepath"
)

const (
	Folder   string = "unixcommand"
	FileName string = "popular-names.txt"
)

func Execute() {
	CountLines()
	ReplaceTabsWithSpace()
	SaveByColumn()
	MergeCol1AndCol2()
	OutputTheFirstNLines(5) // 引数の数値は任意
	OutputTheLastNLines(3)  // 引数の数値は任意
	DivideFileIntoN(500)    // 引数の数値は任意
	DifferentString(filepath.Join(Folder, FileName))
	SortRowsInDescendingOrder()
}

func getFileElements(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// ファイルの中身を全て取得する
	scanner := bufio.NewScanner(f)
	var fileElements []string
	for scanner.Scan() {
		fileElements = append(fileElements, scanner.Text())
	}
	return fileElements, nil
}
