package doc

// Pos is an element of a position identifier. A position identifier identifies an
// atom within a Doc.
type Pos struct {
	Ident uint16
	Site  uint8
}

// ComparePos compares two position identifiers, returning -1 if the left is less than the
// right, 0 if equal, and 1 if greater.
func ComparePos(lp []Pos, rp []Pos) int8 {
	for i := 0; i < len(lp); i++ {
		if len(rp) == i {
			return 1
		}
		if lp[i].Ident < rp[i].Ident {
			return -1
		}
		if lp[i].Ident > rp[i].Ident {
			return 1
		}
		if lp[i].Site < rp[i].Site {
			return -1
		}
		if lp[i].Site > rp[i].Site {
			return 1
		}
	}
	if len(rp) > len(lp) {
		return -1
	}
	return 0
}

// GeneratePos generates a new position identifier between the two positions provided.
func GeneratePos(lp []Pos, rp []Pos, site uint8) []Pos {
	return nil
}
