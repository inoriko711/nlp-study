package unixcommand

const (
	Folder   string = "unixcommand"
	FileName string = "popular-names.txt"
)

func Execute() {
	CountLines()
	ReplaceTabsWithSpace()
	SaveByColumn()
	MergeCol1AndCol2()
	OutputTheFirstNLines(5) // 引数の数値は任意
	OutputTheLastNLines(3)  // 引数の数値は任意
	DivideFileIntoN(500)    // 引数の数値は任意
}
