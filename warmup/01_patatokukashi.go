package warmup

import "fmt"

func Patatokukashi(str string) {
	rs := []rune(str)
	var rs1, rs2 []rune
	for i, r := range rs {
		if i%2 == 0 {
			rs1 = append(rs1, r)
		} else {
			rs2 = append(rs2, r)
		}
	}
	fmt.Println("1.01. " + string(rs1))
	fmt.Println("1.01. " + string(rs2))
}
