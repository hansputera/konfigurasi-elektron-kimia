package functions

import (
	"math"
	"tugas-kimia/types"
)

func findLengthItemInGroup(item *types.SubKulit, groups *[][]*types.SubKulit) int {
	count := 0
	for _, group := range *groups {
		for _, node := range group {
			if node == item {
				count++
			}
		}
	}

	return count
}

func SubkulitGrouping(plates []*types.SubKulit) *[][]*types.SubKulit {
	subkulits := [][]*types.SubKulit{}

	subkulits = append(subkulits, []*types.SubKulit{plates[0], plates[1]}, []*types.SubKulit{plates[2], plates[3]})
	for index, p := range plates {
		if index > 3 {
			latestSubs := subkulits[len(subkulits)-1]
			if len(latestSubs) == 2 {
				subkulits = append(subkulits, []*types.SubKulit{p})
			} else {
				var node *types.SubKulit
				lenItem := findLengthItemInGroup(p, &subkulits)
				if lenItem != 0 {
					node = plates[lenItem+1]
				} else {
					node = plates[index+2]
				}
				subkulits[len(subkulits)-1] = append(subkulits[len(subkulits)-1], node)

			}
		}
	}

	return &subkulits
}

func SubkulitAtom(atom *types.Atom, templates *[]*types.SubKulit) *[]*types.SubKulit {
	subkulits := []*types.SubKulit{}
	total := 0 // total penjumlahan elektron

	for index, t := range *templates {
		if total >= atom.Nomor {
			break
		}

		sisa := math.Max(float64(total), float64(atom.Nomor)) - math.Min(float64(total), float64(atom.Nomor))
		if sisa < float64(t.Electron) {
			total += int(sisa)

			t.Electron = int(sisa)
			subkulits = append(subkulits, t)
			break
		}

		if index <= 1 {
			total += t.Electron
			subkulits = append(subkulits, t)
		} else {
			subkulitBefore := subkulits[len(subkulits)-1]
			if subkulitBefore.Group != t.Group && subkulitBefore.Electron != t.Electron {
				total += t.Electron
				subkulits = append(subkulits, t)
			}
		}

		// TODO: penggunaan karakter menyamping
		// Dari "2p^6" ke "3s^2"

		// total += t.Electron
		// subkulits = append(subkulits, t)
	}
	return &subkulits
}
