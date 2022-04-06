package utility

import (
	"fmt"
)

type NPMParts struct {
	Year        string
	Departement string
	Division    string
	AbsentCode  string
}

func GenerateNPMRangeList(start, to string) ([]string, error) {
	npm1 := NPMParts{
		Year:        start[0:2],
		Departement: start[2:4],
		Division:    start[4:7],
		AbsentCode:  start[7:10],
	}

	npm2 := NPMParts{
		Year:        to[0:2],
		Departement: to[2:4],
		Division:    to[4:7],
		AbsentCode:  to[7:10],
	}

	fmt.Println(npm1)
	fmt.Println(npm2)

	return []string{""}, nil
}
