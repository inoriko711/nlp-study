package unixcommand

import "testing"

func TestFrequency(t *testing.T) {
	t.Run("TestFrequency", func(t *testing.T) {
		if !Frequency() {
			t.Fatal()
		}
	})
}
