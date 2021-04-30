package regex

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// https://nlp100.github.io/ja/ch03.html#20-json%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E8%AA%AD%E3%81%BF%E8%BE%BC%E3%81%BF
func ReadJsonData() {
	file, err := os.Open(JawikiCoountryJsonGZ)
	if err != nil {
		return
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	defer gzipReader.Close()
	output := bytes.Buffer{}
	output.ReadFrom(gzipReader)
	s := string(output.Bytes())

	if _, err := io.Copy(file, gzipReader); err != nil {
		return
	}

	fmt.Println(gzipReader.Name)
	fmt.Println(s)
}
