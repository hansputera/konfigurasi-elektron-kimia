package types

import "fmt"

type SubKulit struct {
	Identifier int    // 2
	Group      string // s OR p OR d OR f
	Electron   int    // electrons count
	Floor      int
}

func (subkulit *SubKulit) ToString() string {
	return fmt.Sprintf("%d%s^%d", subkulit.Identifier, subkulit.Group, subkulit.Electron)
}
