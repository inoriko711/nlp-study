package warmup

import "fmt"

func Patatokukashi2(str1, str2 string) {
	rs1 := []rune(str1)
	rs2 := []rune(str2)
	var respList []rune
	for i := 0; i < len(rs1); i++ {
		respList = append(respList, rs1[i])
		respList = append(respList, rs2[i])
	}
	fmt.Println("1.02. " + string(respList))
}
