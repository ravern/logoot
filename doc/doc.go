package doc

// Doc represents a Logoot document. Actions like Insert and Delete can be performed on Doc.
// If at any time an invalid position is given, a panic will occur, so raw positions should
// only be used for debugging purposes.
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
	for _, c := range content {
		// End will always exist.
		d.InsertLeft(End, c)
	}
	d.Insert(End, "")
	return d
}

/* Basic methods */

// Get the atom at the position. Secondary return value indicates whether the value exists.
func (d *Doc) Get(p []Pos) (string, bool) {
	pr := d.pairs
	for {
		if len(pr) == 0 {
			// No position found
			return "", false
		}
		spt := len(pr) / 2
		pair := pr[spt]
		if cmp := ComparePos(pair.pos, p); cmp == 0 {
			return pair.atom, true
		} else if cmp == -1 {
			pr = pr[spt+1:]
		} else if cmp == 1 {
			pr = pr[0:spt]
		}
	}
}

// Insert a new pair at the position, returning success or failure (already existing
// position).
func (d *Doc) Insert(p []Pos, atom string) bool {
	d.pairs = append(d.pairs, pair{p, atom})
	return false
}

// Delete the pair at the position, returning success or failure (non-existent position).
func (d *Doc) Delete(p []Pos) bool {
	return false
}

// Left returns the position to the left of the given position, and a flag indicating
// whether it exists (when the given position is the start, there is no position to the
// left of it).
func (d *Doc) Left(p []Pos) ([]Pos, bool) {
	return nil, false
}

// GeneratePos generates a new position identifier between the two positions provided.
func (d *Doc) GeneratePos(lp []Pos, rp []Pos) []Pos {
	return nil
}

// Right returns the position to the right of the given position, and a flag indicating
// whether it exists (when the given position is the end, there is no position to the
// right of it).
func (d *Doc) Right(p []Pos) ([]Pos, bool) {
	return nil, false
}

// Each element, along with their positions will be passed to the provided closure in order.
// The closure should return true to continue or false to break.
func (d *Doc) Each(f func([]Pos, string) bool) {
}

/* Convenience methods */

// InsertLeft inserts the atom to the left of the given position, returning whether it is
// successful (when the given position doesn't exist, InsertLeft won't do anything and
// return false).
func (d *Doc) InsertLeft(p []Pos, atom string) bool {
	return false
}

// InsertRight inserts the atom to the right of the given position, returning whether it is
// successful (when the given position doesn't exist, InsertRight won't do anything and
// return false).
func (d *Doc) InsertRight(p []Pos, atom string) bool {
	return false
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

// Content of the entire document, except the first and last empty lines.
func (d *Doc) Content() []string {
	return nil
}
