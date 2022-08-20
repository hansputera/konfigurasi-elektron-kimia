package parsers

import (
	"regexp"
	"strings"
	"tugas-kimia/types"
)

func ParseSubKulitStr(raw *string) *[]*types.SubKulit {
	subkulits := []*types.SubKulit{}
	regex := regexp.MustCompile("[0-9]+(s|p|d|f)" + regexp.QuoteMeta("^") + "[0-9]+")
	res := strings.Split((*raw), "\n")
	floor := 1

	for index, r := range res {
		if regex.MatchString(r) {
			items := regex.FindAllString(r, 10)
			for _, item := range items {
				subkulits = append(subkulits, &types.SubKulit{
					Identifier: item,
					Floor:      floor,
				})
			}
			floor++
		} else {
			res[index] = ""
		}
	}

	return &subkulits
}
