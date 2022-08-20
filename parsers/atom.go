package parsers

import (
	"strconv"
	"strings"
	"tugas-kimia/types"
)

func GetAtom(text string) *types.Atom {
	// 22.9898^11^Na
	// {22.9898 11 Na}
	// A = 22.9898; Z = 11; X = Na

	A := 0.0
	Z := 0
	X := ""

	splits := strings.Split(text, "^")

	X = splits[len(splits)-1]
	if len(splits) == 3 {
		A, _ = strconv.ParseFloat(strings.ReplaceAll(splits[0], ",", "."), 32)
		Z, _ = strconv.Atoi(splits[1])
	} else {
		Z, _ = strconv.Atoi(splits[0])
	}

	return &types.Atom{
		Unsur: X,
		Nomor: Z,
		Massa: A,
	}
}
