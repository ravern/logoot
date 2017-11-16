package doc_test

import (
	"strings"
	"testing"

	"github.com/ravernkoh/logoot/doc"
)

// Can only test start and end positions because the rest are randomly generated.
func TestDocGetPos(t *testing.T) {
	tests := []struct {
		pos    []doc.Pos
		atom   string
		exists bool
	}{
		{
			doc.Start,
			"",
			true,
		},
		{
			doc.End,
			"",
			true,
		},
	}

	for i, test := range tests {
		d := doc.New(strings.Split("hello world", ""), 1)
		atom, exists := d.Get(test.pos)

		if atom != test.atom || exists != test.exists {
			t.Errorf("Test %d: Expected %s, %t, got %s, %t.", i+1, test.atom, test.exists, atom, exists)
		}
	}
}
