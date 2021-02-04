package unixcommand

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func OutputTheLastNLines(n int) {
	result := checkOutputTheLastNLines(n, nil)
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

	fmt.Println(want)
	// return reflect.DeepEqual(want, got)
	return false
}
