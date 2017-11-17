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

func TestDocInsertAndContent(t *testing.T) {
	tests := []struct {
		poss    [][]doc.Pos
		atoms   []string
		content string
	}{
		{
			[][]doc.Pos{
				[]doc.Pos{{1, 1}},
				[]doc.Pos{{3, 1}},
				[]doc.Pos{{4, 1}},
				[]doc.Pos{{6, 1}},
			},
			[]string{
				"hel",
				"lo ",
				"wor",
				"ld",
			},
			"hello world",
		},
		{
			[][]doc.Pos{
				[]doc.Pos{{1, 1}},
				[]doc.Pos{{1, 1}, {5, 1}},
				[]doc.Pos{{1, 1}, {5, 1}, {1, 1}},
				[]doc.Pos{{2, 1}},
			},
			[]string{
				"boo ",
				"lmao ",
				"test ",
				"now",
			},
			"boo lmao test now",
		},
	}

	for i, test := range tests {
		d := doc.New([]string{}, 1)
		for i, pos := range test.poss {
			d.Insert(pos, test.atoms[i])
		}

		c := d.Content()
		if c != test.content {
			t.Errorf("Test %d: Expected %s, got %s.", i+1, test.content, c)
		}
	}
}

func TestDocLeft(t *testing.T) {
	tests := []string{
		"hel",
		"lo ",
		"wor",
		"ld.",
	}

	d := doc.New(tests, 1)
	p := doc.End
	for i := 0; i < len(tests); i++ {
		np, exists := d.Left(p)
		if !exists && i != len(tests)-1 {
			t.Errorf("Test %d: Expected position to exist.", i+1)
		}
		if doc.ComparePos(np, p) != -1 {
			t.Errorf("Test %d: Expected new position to be less than old position.", i+1)
		}
		p = np
	}
}

func TestDocRight(t *testing.T) {
	tests := []string{
		"tes",
		"tin",
		"g 1",
		"23.",
	}

	d := doc.New(tests, 1)
	p := doc.Start
	for i := 0; i < len(tests); i++ {
		np, exists := d.Right(p)
		if !exists && i != len(tests)-1 {
			t.Errorf("Test %d: Expected position to exist.", i+1)
		}
		if doc.ComparePos(np, p) != 1 {
			t.Errorf("Test %d: Expected new position to be greater than old position.", i+1)
		}
		p = np
	}
}

func TestDocInsertLeft(t *testing.T) {
	tests := []string{
		"hel",
		"lo ",
		"wor",
		"ld.",
	}

	d := doc.New([]string{}, 1)
	p := doc.End
	for i := len(tests) - 1; i >= 0; i-- {
		np, success := d.InsertLeft(p, tests[i])
		if !success {
			t.Errorf("Test %d: Expected successful insert.", i+1)
		}
		p = np
	}

	c := d.Content()
	s := strings.Join(tests, "")
	if c != s {
		t.Errorf("Expected content to be %s, but got %s.", s, c)
	}
}

func TestDocInsertRight(t *testing.T) {
	tests := []string{
		"hel",
		"lo ",
		"wor",
		"ld.",
	}

	d := doc.New([]string{}, 1)
	p := doc.Start
	for i := 0; i < len(tests); i++ {
		np, success := d.InsertRight(p, tests[i])
		if !success {
			t.Errorf("Test %d: Expected successful insert.", i+1)
		}
		p = np
	}

	c := d.Content()
	s := strings.Join(tests, "")
	if c != s {
		t.Errorf("Expected content to be %s, but got %s.", s, c)
	}
}
