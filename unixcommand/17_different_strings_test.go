package unixcommand

import (
	"testing"
)

func TestDifferentString(t *testing.T) {
	UnixcommandFolder = "../unixcommand"
	result := DifferentString()
	if !result {
		t.Fail()
	}
}
