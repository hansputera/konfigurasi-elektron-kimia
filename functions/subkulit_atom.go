package functions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"tugas-kimia/types"
)

func getElektronFromTemplate(template *string) *int {
	elektron, _ := strconv.Atoi(strings.Split(*template, "^")[1])
	return &elektron
}

func SubkulitAtom(atom *types.Atom, templates *[]*types.SubKulit) *[]string {
	subkulits := []string{}
	total := 0 // total penjumlahan elektron

	for _, t := range *templates {
		if total >= atom.Nomor {
			break
		}

		elektron := getElektronFromTemplate(&t.Identifier)
		sisa := math.Max(float64(total), float64(atom.Nomor)) - math.Min(float64(total), float64(atom.Nomor))
		if sisa < float64(*elektron) {
			total += int(sisa)
			subkulits = append(subkulits, strings.Replace(t.Identifier, strconv.Itoa(*elektron), strconv.Itoa(int(sisa)), 1))
			break
		}

		// TODO: penggunaan karakter menyamping
		// Dari "2p^6" ke "3s^2"

		total += *elektron
		subkulits = append(subkulits, t.Identifier)
	}

	fmt.Println(subkulits)
	return &subkulits
}
