package functions

import (
	"testing"
	"tugas-kimia/parsers"
)

func TestSubkulitAtom(t *testing.T) {
	atom := parsers.GetAtom("5^B")
	raw := `
		1s^2
		2s^2    2p^6
		3s^2    3p^6    3d^10
		4s^2    4p^6    4d^10   4f^14
		5s^2    5p^6    5d^10
		6s^2    6p^6 
		7s^2
	`

	subkulit := parsers.ParseSubKulitStr(&raw)
	subkulits := SubkulitAtom(atom, subkulit)

	t.Log(subkulits)
}
