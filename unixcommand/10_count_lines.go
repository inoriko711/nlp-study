package unixcommand

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// % cat unixcommand/popular-names.txt | wc -l
// 	2780
func CountLines() {
	f, err := os.Open(filepath.Join(UnixcommandFolder, PopularNamesFileName))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	lines := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines += 1
	}

	fmt.Printf("2.10. Read %d lines\n", lines)
}
