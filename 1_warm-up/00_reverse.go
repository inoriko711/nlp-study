package main

import "fmt"

func Reverse() {
	str := "stressed"
	rs := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	fmt.Println(string(rs))
}
