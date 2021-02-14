package unixcommand

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func SortRowsInDescendingOrder() *[]string {

	result := checkSortRowsInDescendingOrder()

	fmt.Printf("2.18. %t\n", result)
	return nil
}

// 答え合わせ用
func checkSortRowsInDescendingOrder() bool {
	// % sort -n -k 3 unixcommand/popular-names.txt
	out, err := exec.Command("sort", "-n", "-k", "3", filepath.Join(Folder, FileName)).Output()
	if err != nil {
		log.Println(err)
		return false
	}

	wants := strings.Split(string(out), "\n")

	fmt.Println(wants)
	return false
}
