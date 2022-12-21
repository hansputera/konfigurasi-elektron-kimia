package parsers

import "testing"

func TestParseSubkulit(t *testing.T) {
	raw := ReadSubkulitFile("../assets/subkulit.txt")

	subkulits := ParseSubKulitSchema(&raw)
	for _, kulit := range *subkulits {
		t.Log("Id:", kulit.Identifier, "| G:", kulit.Group, "| E:", kulit.Electron)
	}
}
