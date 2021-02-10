package unixcommand

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

// DifferentString 1列目の文字列の種類（異なる文字列の集合）を求める関数
func DifferentString() []string {

	// TODO goでの実装

	fmt.Println("2.17. 1列目の文字列の種類：")

	result := checkDifferentString()
	fmt.Printf("2.17. %t\n", result)

	return nil
}

// TODO 実装
func checkDifferentString() bool {
	// % cut -f 1 unixcommand/popular-names.txt | sort | uniq
	cmdstr := fmt.Sprintf("cut -f 1 %s | sort | uniq", filepath.Join(Folder, FileName))

	out, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
		return false
	}
	outList := strings.Split(string(out), "\n")

	fmt.Println(outList)
	return false
}
