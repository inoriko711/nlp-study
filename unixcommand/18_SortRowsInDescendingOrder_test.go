package unixcommand

import "testing"

func TestSortRowsInDescendingOrder(t *testing.T) {
	// % sort -n -k 3 unixcommand/popular-names.txt
	t.Run("TestSortRowsInDescendingOrder", func(t *testing.T) {
		if !SortRowsInDescendingOrder() {
			t.Fatal()
		}
	})
}
