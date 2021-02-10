package unixcommand

import (
	"fmt"
	"os/exec"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestDifferentString(t *testing.T) {
	// golangで準備した結果を取得
	got := DifferentString()

	// 実際のコマンドを準備'cut -f 1 unixcommand/popular-names.txt | sort | uniq'
	cmdstr := fmt.Sprintf("cut -f 1 %s | sort | uniq", FileName)

	// 実行結果の取得
	out, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		t.Fatal(err)
	}
	want := strings.Split(string(out), "\n")

	sort.Strings(want)
	want = revoke(want, "")

	// 結果が相違ないことの確認
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v\n got %v", want, got)
	}

}
