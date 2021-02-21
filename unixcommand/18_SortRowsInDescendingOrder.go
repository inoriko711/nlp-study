package unixcommand

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/kylelemons/godebug/pretty"
)

// https://nlp100.github.io/ja/ch02.html#18-%E5%90%84%E8%A1%8C%E3%82%923%E3%82%B3%E3%83%A9%E3%83%A0%E7%9B%AE%E3%81%AE%E6%95%B0%E5%80%A4%E3%81%AE%E9%99%8D%E9%A0%86%E3%81%AB%E3%82%BD%E3%83%BC%E3%83%88
func SortRowsInDescendingOrder() bool {
	type rowByPeople struct {
		People int
		row    string
	}

	// popular-names.txtを開く
	f, err := os.Open(filepath.Join(UnixcommandFolder, PopularNamesFileName))
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer f.Close()

	rowByPeopleList := make([]rowByPeople, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, "\t")

		// 3列目をint型で取得
		key, err := strconv.Atoi(words[2])
		if err != nil {
			log.Fatal(err)
			return false
		}

		rowByPeopleList = append(rowByPeopleList, rowByPeople{
			People: key,
			row:    line,
		})
	}

	sort.Slice(rowByPeopleList, func(i, j int) bool { return rowByPeopleList[i].People > rowByPeopleList[j].People })

	var resList []string
	for _, r := range rowByPeopleList {
		resList = append(resList, r.row)
	}

	err = checkSortRowsInDescendingOrder(resList)
	if err != nil {
		log.Print(err)
		return false
	}

	// 結果の出力
	// fmt.Println(resList)

	fmt.Println("2.18. ok")
	return true
}

// 答え合わせ用
func checkSortRowsInDescendingOrder(gotRows []string) error {
	// % sort -n -k 3 -r unixcommand/popular-names.txt
	out, err := exec.Command("sort", "-n", "-k", "3", "-r", filepath.Join(UnixcommandFolder, PopularNamesFileName)).Output()
	if err != nil {
		return err
	}

	wantRows := strings.Split(string(out), "\n")
	wantRows = revoke(wantRows, "")
	gotRows = revoke(gotRows, "")

	// この問題に関しては人数の値の並びだけ一致してればいい
	want := getPeopleColumn(wantRows)
	got := getPeopleColumn(gotRows)
	if !reflect.DeepEqual(want, got) {
		return errors.New(pretty.Compare(want, got))
	}
	return nil
}

func getPeopleColumn(list []string) []string {
	var res []string
	for _, element := range list {
		columns := strings.Split(string(element), "\t")
		res = append(res, columns[2])
	}
	return res
}
