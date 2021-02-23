package unixcommand

import (
	"errors"
	"fmt"
	"log"
)

// 各行の1列目の文字列の出現頻度を求め，その高い順に並べて表示せよ．
// 確認にはcut, uniq, sortコマンドを用いよ．
func Frequency() bool {
	// TODO 実装する

	err := checkFrequency()
	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println("2.19 ok")
	return false
}

func checkFrequency() error {
	// TODO 実装する
	return errors.New("未実装")
}
