package functions

import (
	"testing"
	"tugas-kimia/parsers"
)

func TestKulitAtom(t *testing.T) {
	// Test 54^Xe
	atom := parsers.GetAtom("54^Xe")
	kulits := KulitAtom(atom)

	// 54^Xe = 2 8 18 18 8
	if kulits[0] != 2 || kulits[1] != 8 || kulits[2] != 18 || kulits[3] != 18 || kulits[4] != 8 {
		t.Fail()
	}

	// Test 10^Ne
	atom = parsers.GetAtom("10^Ne")
	kulits = KulitAtom(atom)

	// 10^Ne = 2 8
	if kulits[0] != 2 || kulits[1] != 8 {
		t.Fail()
	}

	// Test 88^Ra
	atom = parsers.GetAtom("88^Ra")
	kulits = KulitAtom(atom)

	// 88^Ra = 2 8 18 32 18 8 2
	if kulits[0] != 2 || kulits[1] != 8 || kulits[2] != 18 || kulits[3] != 32 || kulits[4] != 18 || kulits[5] != 8 || kulits[6] != 2 {
		t.Fail()
	}

	// Test 19^K
	atom = parsers.GetAtom("19^K")
	kulits = KulitAtom(atom)

	// 19^K = 2 8 8 1
	if kulits[0] != 2 || kulits[1] != 8 || kulits[2] != 8 || kulits[3] != 1 {
		t.Fail()
	}
}
