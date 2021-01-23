package unixcommand

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ReplaceTabsWithSpaces() {
	f, err := os.Open(filepath.Join("unixcommand", "popular-names.txt"))
	if err != nil {
		log.Fatal(err)
		return
	}

	var cal1Word []string
	// var cal2Word []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		cal1Word = append(cal1Word, line)
	}

	fmt.Println(cal1Word)

	// 各行の1列目だけを抜き出したもの
	col1File, err := os.OpenFile(filepath.Join("", ""), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer col1File.Close()

	// 各行の2列目だけを抜き出したもの
	col2File, err := os.OpenFile(filepath.Join("", ""), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer col2File.Close()
}
