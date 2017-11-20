package doc_test

import (
	"fmt"
	"testing"

	"github.com/ravernkoh/logoot/doc"
)

func TestComparePos(t *testing.T) {
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

func TestGeneratePos(t *testing.T) {
	tests := []struct {
		lpos    []doc.Pos
		rpos    []doc.Pos
		success bool
	}{
		{
			[]doc.Pos{{0, 0}},
			[]doc.Pos{{1, 1}},
			true,
		},
		{
			[]doc.Pos{{1, 1}},
			[]doc.Pos{{1, 1}, {3, 2}},
			true,
		},
		{
			[]doc.Pos{{5, 3}, {11, 10}},
			[]doc.Pos{{5, 3}, {11, 10}},
			false,
		},
		{
			[]doc.Pos{{1, 1}, {5, 4}},
			[]doc.Pos{{1, 1}, {3, 2}},
			false,
		},
		{
			[]doc.Pos{{1, 1}, {5, 4}},
			[]doc.Pos{{1, 1}, {5, 4}, {3, 2}},
			true,
		},
		{
			[]doc.Pos{{0, 1}, {3, 4}},
			[]doc.Pos{{1, 1}},
			true,
		},
		{
			[]doc.Pos{{0, 1}, {65534, 1}},
			[]doc.Pos{{1, 1}},
			true,
		},
	}

	for i, test := range tests {
		p, success := doc.GeneratePos(test.lpos, test.rpos, 1)

		l := doc.ComparePos(test.lpos, p)
		r := doc.ComparePos(p, test.rpos)
		if success && (l != -1 || r != -1) {
			fmt.Printf("Test %d, Pos %v\n", i+1, p)
			t.Errorf("Test %d: Expected generated Pos to be between given Pos(s).", i+1)
		}
		if success != test.success {
			t.Errorf("Test %d: Expected success flag to be %t, but got %t", i+1, test.success, success)
		}
	}
}

func TestPosByteConversion(t *testing.T) {
	tests := [][]doc.Pos{
		[]doc.Pos{{0, 0}},
		[]doc.Pos{{1, 1}},
		[]doc.Pos{{1, 1}},
		[]doc.Pos{{1, 1}},
		[]doc.Pos{{1, 1}, {3, 2}},
		[]doc.Pos{{5, 3}, {11, 10}},
		[]doc.Pos{{5, 3}, {11, 10}},
		[]doc.Pos{{1, 1}, {5, 4}},
		[]doc.Pos{{0, 1}, {65534, 1}},
		[]doc.Pos{{1, 1}, {3, 2}},
		[]doc.Pos{{1, 1}, {5, 4}},
		[]doc.Pos{{1, 1}, {5, 4}, {3, 2}},
		[]doc.Pos{{0, 1}, {3, 4}},
		[]doc.Pos{{1, 1}},
	}

	for i, test := range tests {
		b := doc.PosBytes(test)
		p := doc.NewPos(b)
		if doc.ComparePos(p, test) != 0 {
			fmt.Println("")
			fmt.Println(test)
			fmt.Println(p)
			t.Errorf("Test %d: Expected converting back and forth to return same pos.", i+1)
		}
	}
}
