package doc

import (
	"math/rand"
)

// Pos is an element of a position identifier. A position identifier identifies an
// atom within a Doc. The behaviour of an empty position identifier (length == 0) is
// undefined, so just do not pass in empty position identifiers to any method/function.
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

// random number between x and y, where y is greater than x.
func random(x, y uint16) uint16 {
	return uint16(rand.Intn(int(y-x-1))) + 1 + x
}

// GeneratePos generates a new position identifier between the two positions provided.
// Secondary return value indicates whether it was successful (when the two positions
// are equal, or the left is greater than right, position cannot be generated).
func GeneratePos(lp, rp []Pos, site uint8) ([]Pos, bool) {
	if ComparePos(lp, rp) != -1 {
		return nil, false
	}
	p := []Pos{}
	for i := 0; i < len(lp); i++ {
		l := lp[i]
		r := rp[i]
		if l.Ident == r.Ident && l.Site == r.Site {
			p = append(p, Pos{l.Ident, l.Site})
			continue
		}
		if d := r.Ident - l.Ident; d > 1 {
			r := random(l.Ident, r.Ident)
			p = append(p, Pos{r, site})
		} else if d == 1 {
			if site > l.Site {
				p = append(p, Pos{l.Ident, site})
			} else if site < r.Site {
				p = append(p, Pos{r.Ident, site})
			} else {
				r := random(0, ^uint16(0))
				p = append(p, Pos{l.Ident, l.Site}, Pos{r, site})
			}
		} else {
			if site > l.Site && site < r.Site {
				p = append(p, Pos{l.Ident, site})
			} else {
				r := random(0, ^uint16(0))
				p = append(p, Pos{l.Ident, l.Site}, Pos{r, site})
			}
		}
		return p, true
	}
	if len(rp) > len(lp) {
		r := random(0, rp[len(lp)].Ident)
		p = append(p, Pos{r, site})
	}
	return p, true
}
