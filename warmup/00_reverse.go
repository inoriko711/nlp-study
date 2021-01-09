package warmup

import "fmt"

func Reverse(str string) {
	rs := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	fmt.Println("1.00. " + string(rs))
}
