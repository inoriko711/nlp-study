package regex

import (
	"fmt"
	"os"
)

var JawikiCoountryJsonGZ string

func init() {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	JawikiCoountryJsonGZ = fmt.Sprintf("%s/regex/jawiki-country.json.gz", dir)
}

func Execute() {
}
