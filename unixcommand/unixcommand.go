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
	OutputTheFirstNLines(5)
	OutputTheLastNLines(3)
}
