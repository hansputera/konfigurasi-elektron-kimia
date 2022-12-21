package main

import (
	"fmt"
	"tugas-kimia/functions"
	"tugas-kimia/parsers"
)

func main() {
	atom := parsers.GetAtom("36^Kr")

	// konfigurasi elektron
	kulitAtom := functions.KulitAtom(atom)
	fmt.Println("Kulit Atom (14Cl):", kulitAtom)

	// subkulit
	schema := parsers.ReadSubkulitFile("assets/subkulit.txt")
	subkulits := parsers.ParseSubKulitSchema(&schema)
	groupings := functions.SubkulitGrouping(*subkulits)

	fmt.Println((*subkulits)[0])
	for index, g := range *groupings {
		for _, x := range g {
			fmt.Println("Kelompok", index+1, "punya", x.ToString())
		}
	}
}
