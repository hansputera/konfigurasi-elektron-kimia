package functions

import (
	"tugas-kimia/types"
)

// rumus: 2n^2
func maxElektronByPos(pos int) int {
	return 2 * ((pos + 1) * (pos + 1))
}

func KulitAtom(atom *types.Atom) *[7]int {
	// K, L, M, N, O, P, Q
	kulits := [7]int{0, 0, 0, 0, 0, 0, 0}
	pos := 0
	total := 0
	nomor := atom.Nomor

	for nomor > 0 && pos < len(kulits) {
		max_elektron := maxElektronByPos(pos)
		if max_elektron > 32 {
			max_elektron = 32
		}
		sisa := nomor - total

		if sisa > max_elektron || sisa == max_elektron {
			kulits[pos] = max_elektron
		} else {
			if (total+sisa) == atom.Nomor && (total+sisa) < max_elektron {
				kulits[pos] = sisa
				break
			}
			max_elektron = maxElektronByPos(pos - 1)
			if sisa > max_elektron || sisa == max_elektron {
				kulits[pos] = max_elektron
			} else {
				posCopy := pos
				for sisa > 0 {
					max_elektron = maxElektronByPos(posCopy - 1)
					if sisa > max_elektron || sisa == max_elektron {
						kulits[pos] = max_elektron
						sisa -= kulits[posCopy]
					} else {
						if posCopy <= 2 {
							kulits[pos] = sisa
							break
						}
						posCopy--
					}
				}
			}
			// poscpy := pos
			// for sisa > 0 {
			// 	max_elektron = maxElektronByPos(poscpy - 1)
			// 	if sisa > max_elektron || sisa == max_elektron {
			// 		kulits[pos] = max_elektron
			// 		sisa -= kulits[poscpy]
			// 	} else {
			// 		if poscpy == 1 && sisa <= 2 && sisa > 0 {
			// 			kulits[pos] = sisa
			// 			break
			// 		}
			// 		poscpy--
			// 	}

			// 	total += kulits[pos]
			// }
		}

		total += kulits[pos]
		pos++
	}

	return &kulits
}
