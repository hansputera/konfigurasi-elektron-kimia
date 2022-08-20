package main

import (
	"fmt"
	"tugas-kimia/functions"
	"tugas-kimia/parsers"
)

func main() {
	atom := parsers.GetAtom("19^K")

	// konfigurasi elektron
	kulitAtom := functions.KulitAtom(atom)
	fmt.Println(kulitAtom)
}
