package parsers

import (
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"tugas-kimia/types"
)

func ReadSubkulitFile(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln(err.Error())
		return ""
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err.Error())
		return ""
	}

	return string(data)
}

func toInt(str string) int {
	integer, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return integer
}

func ParseSubKulit(str string) *types.SubKulit {
	regex := regexp.MustCompile("([0-9]+)(s|p|d|f)" + regexp.QuoteMeta("^") + "([0-9]+)")
	if !regex.MatchString(str) {
		return nil
	}
	matches := regex.FindAllStringSubmatch(str, -1)
	if len(matches) < 1 || len(matches[0]) < 3 {
		return nil
	}

	return &types.SubKulit{
		Identifier: toInt(matches[0][1]),
		Group:      matches[0][2],
		Electron:   toInt(matches[0][3]),
	}
}

func ParseSubKulitSchema(raw *string) *[]*types.SubKulit {
	subkulits := []*types.SubKulit{}
	regex := regexp.MustCompile("[0-9]+(s|p|d|f)" + regexp.QuoteMeta("^") + "[0-9]+")
	res := strings.Split((*raw), "\n")
	// floor := 1

	for index, floor := range res {
		if regex.MatchString(floor) {
			items := regex.FindAllString(floor, 10)
			for _, item := range items {
				subkulit := ParseSubKulit(item)

				subkulit.Floor = index + 1
				subkulits = append(subkulits, subkulit)
			}
		} else {
			res[index] = ""
		}
	}

	// for index, r := range res {
	// 	if regex.MatchString(r) {
	// 		items := regex.FindAllString(r, 10)
	// 		for _, item := range items {
	// 			subkulits = append(subkulits, &types.SubKulit{
	// 				Identifier: item,
	// 				Floor:      floor,
	// 			})
	// 		}
	// 		floor++
	// 	} else {
	// 		res[index] = ""
	// 	}
	// }

	return &subkulits
}
