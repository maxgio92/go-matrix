package main

import "fmt"

type combinazione string
type part struct {
	pointer      int
	combinazioni []combinazione
}

var (
	parts = []*part{
		&part{combinazioni: []combinazione{"A", "B"}, pointer: 0},      // part
		&part{combinazioni: []combinazione{"1", "2", "3"}, pointer: 0}, // part
		&part{combinazioni: []combinazione{"x", "y", "z"}, pointer: 0}, // part
	}
)

// (i)	(j)
// A	1	Z
// B	2	Y
// C		X
//			W

// n 			= the column index
// part.pointer	= the row index
func main() {
	combinazioni := []*string{}
	combinazione_cumulata := ""
	y := 0

	for n, _ := range parts {
		// TODO: it should start from scratch
		combinazione_cumulata = ""
		n = n

		gotoNextPart(&combinazioni, &combinazione_cumulata, 0, parts[0], parts, y)
	}
}

func gotoNextPart(combinazioni *[]*string, combinazione_cumulata *string, n int, part *part, parts []*part, y int) {

	fmt.Println(*part)

	if n+1 < len(parts) {
		fmt.Println("avanti")

		*combinazione_cumulata += string(part.combinazioni[part.pointer])

		n++
		part = parts[n]
		gotoNextPart(combinazioni, combinazione_cumulata, n, part, parts, y)
	} else {
		fmt.Println("indietro, sono all'ultima parte")

		for _, combinazione := range part.combinazioni {
			fmt.Println(*combinazione_cumulata + string(combinazione))

			tmp := string(*combinazione_cumulata + string(combinazione))
			*combinazioni = append(*combinazioni, &tmp)
		}

		n--
		part = parts[n]

		part.pointer++
	}
}
