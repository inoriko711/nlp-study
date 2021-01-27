package unixcommand

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func MergeCol1AndCol2() {
	//
	out, err := exec.Command("paste", filepath.Join("unixcommand", "col1.txt"), filepath.Join("unixcommand", "col2.txt")).Output()
	if err != nil {
		log.Fatal(err)
	}
	outList := strings.Split(string(out), "\n")
	fmt.Println(outList)
}
