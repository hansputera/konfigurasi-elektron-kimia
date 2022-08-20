package parsers

import "testing"

func TestParseSubkulit(t *testing.T) {
	raw := `

		1s^2
		2s^2    2p^6
		3s^2    3p^6    3d^10
		4s^2    4p^6    4d^10   4f^14
		5s^2    5p^6    5d^10
		6s^2    6p^6 
		7s^2
	`

	subkulits := ParseSubKulitStr(&raw)
	for _, kulit := range *subkulits {
		t.Log("Subkulit", kulit.Identifier, "di lantai", kulit.Floor)
	}
}
