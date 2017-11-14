package doc_test

import (
	"testing"

	"github.com/ravernkoh/logoot/doc"
)

func TestPosCompare(t *testing.T) {
	tests := []struct {
		lpos   []doc.Pos
		rpos   []doc.Pos
		result int8
	}{
		{
			[]doc.Pos{{0, 0}},
			[]doc.Pos{{1, 1}},
			-1,
		},
		{
			[]doc.Pos{{1, 1}},
			[]doc.Pos{{1, 1}, {3, 2}},
			-1,
		},
		{
			[]doc.Pos{{5, 3}, {11, 10}},
			[]doc.Pos{{5, 3}, {11, 10}},
			0,
		},
		{
			[]doc.Pos{{1, 1}, {5, 4}},
			[]doc.Pos{{1, 1}, {3, 2}},
			1,
		},
		{
			[]doc.Pos{{1, 1}, {5, 4}, {3, 2}},
			[]doc.Pos{{1, 1}, {5, 4}},
			1,
		},
	}

	for i, test := range tests {
		res := doc.ComparePos(test.lpos, test.rpos)
		if res != test.result {
			t.Errorf("Test %d: Expected %d, got %d.", i+1, test.result, res)
		}
	}
}
