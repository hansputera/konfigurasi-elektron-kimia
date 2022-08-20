package parsers

import (
	"testing"
)

func Test_GetAtomWithNomor(t *testing.T) {
	atom := GetAtom("4^Be")
	if atom.Unsur != "Be" || atom.Nomor != 4 {
		t.Fail()
	}
}

func Test_GetAtomWithMassa(t *testing.T) {
	atom := GetAtom("22.9898^11^Na")
	if atom.Unsur != "Na" || atom.Nomor != 11 || atom.Massa > 23 || atom.Massa < 22 {
		t.Fail()
	}
}
