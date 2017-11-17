package doc

import "bytes"

// Doc represents a Logoot document. Actions like Insert and Delete can be performed
// on Doc. If at any time an invalid position is given, a panic will occur, so raw
// positions should only be used for debugging purposes.
type Doc struct {
	site  uint8
	pairs []pair
}

// pair is a position identifier and its atom.
type pair struct {
	pos  []Pos
	atom string
}

// Start and end positions. These will always exist within a document.
var (
	Start = []Pos{{0, 0}}
	End   = []Pos{{^uint16(0), 0}}
)

// New creates a new document containing the given content.
func New(content []string, site uint8) *Doc {
	d := &Doc{site: site}
	d.Insert(Start, "")
	d.Insert(End, "")
	for _, c := range content {
		// End will always exist.
		d.InsertLeft(End, c)
	}
	return d
}

/* Basic methods */

// Index of a position in the Doc. Secondary value indicates whether the value exists.
// If the value doesn't exist, the index returned is the index that the position would
// have been in, should it have existed.
func (d *Doc) Index(p []Pos) (int, bool) {
	off := 0
	pr := d.pairs
	for {
		if len(pr) == 0 {
			return off, false
		}
		spt := len(pr) / 2
		pair := pr[spt]
		if cmp := ComparePos(pair.pos, p); cmp == 0 {
			return spt + off, true
		} else if cmp == -1 {
			off += spt + 1
			pr = pr[spt+1:]
		} else if cmp == 1 {
			pr = pr[0:spt]
		}
	}
}

// Atom at the position. Secondary return value indicates whether the value exists.
func (d *Doc) Get(p []Pos) (string, bool) {
	i, exists := d.Index(p)
	if !exists {
		return "", false
	}
	return d.pairs[i].atom, true
}

// Insert a new pair at the position, returning success or failure (already existing
// position).
func (d *Doc) Insert(p []Pos, atom string) bool {
	i, exists := d.Index(p)
	if exists {
		return false
	}
	d.pairs = append(d.pairs[0:i], append([]pair{{p, atom}}, d.pairs[i:]...)...)
	return true
}

// Delete the pair at the position, returning success or failure (non-existent position).
func (d *Doc) Delete(p []Pos) bool {
	return false
}

// Left returns the position to the left of the given position, and a flag indicating
// whether it exists (when the given position is the start, there is no position to the
// left of it). Will be false if the given position is invalid. The Start pair is not
// considered as an actual pair.
func (d *Doc) Left(p []Pos) ([]Pos, bool) {
	i, exists := d.Index(p)
	if !exists || i == 0 {
		return nil, false
	}
	return d.pairs[i-1].pos, true
}

// Right returns the position to the right of the given position, and a flag indicating
// whether it exists (when the given position is the end, there is no position to the
// right of it). Will be false if the given position is invalid. The End pair is not
// considered as an actual pair.
func (d *Doc) Right(p []Pos) ([]Pos, bool) {
	i, exists := d.Index(p)
	if !exists || i >= len(d.pairs)-1 {
		return nil, false
	}
	return d.pairs[i+1].pos, true
}

// GeneratePos generates a new position identifier between the two positions provided.
// Secondary return value indicates whether it was successful (when the two positions
// are equal, or the left is greater than right, position cannot be generated).
func (d *Doc) GeneratePos(lp []Pos, rp []Pos) ([]Pos, bool) {
	return GeneratePos(lp, rp, d.site)
}

/* Convenience methods */

// InsertLeft inserts the atom to the left of the given position, returning whether it
// is successful (when the given position doesn't exist, InsertLeft won't do anything
// and return false).
func (d *Doc) InsertLeft(p []Pos, atom string) bool {
	lp, success := d.Left(p)
	if !success {
		return false
	}
	np, success := d.GeneratePos(lp, p)
	if !success {
		return false
	}
	d.Insert(np, atom)
	return true
}

// InsertRight inserts the atom to the right of the given position, returning whether it
// is successful (when the given position doesn't exist, InsertRight won't do anything
// and return false).
func (d *Doc) InsertRight(p []Pos, atom string) bool {
	rp, success := d.Right(p)
	if !success {
		return false
	}
	np, success := d.GeneratePos(p, rp)
	if !success {
		return false
	}
	d.Insert(np, atom)
	return true
}

// DeleteLeft deletes the atom to the left of the given position, returning whether it
// was successful (when the given position is the start, there is no position to the left
// of it).
func (d *Doc) DeleteLeft(p []Pos, atom string) bool {
	return false
}

// DeleteRight deletes the atom to the right of the given position, returning whether it
// was successful (when the given position is the end, there is no position to the right
// of it).
func (d *Doc) DeleteRight(p []Pos, atom string) bool {
	return false
}

// Content of the entire document.
func (d *Doc) Content() string {
	var b bytes.Buffer
	for i := 1; i < len(d.pairs)-1; i++ {
		b.WriteString(d.pairs[i].atom)
	}
	return b.String()
}
