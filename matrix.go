package main

import "fmt"

type combinazione string
type part struct {
	pointer      int
	combinazioni []combinazione
}

var (
	parts = []part{
		part{combinazioni: []combinazione{"A", "B"}, pointer: 0},                // part
		part{combinazioni: []combinazione{"1", "2", "3", "4", "5"}, pointer: 0}, // part
		part{combinazioni: []combinazione{"w", "x", "y", "z"}, pointer: 0},      // part
		//&part{combinazioni: []combinazione{"w", "x", "y", "z"}, pointer: 0},      // part // TODO: fix this
	}
)

// y	(i)	(j)
// ^	A	1	Z
// |	B	2	Y
// |		3	X
// ------> x

// n 			= the column index
// part.pointer	= the row index
func main() {
	combinazioni := []string{}
	combinazione_cumulata := ""
	y := 0

	// Print
	for _, v := range parts {
		fmt.Println(v)
	}

	for {
		combinazione_cumulata = ""

		gotoNextPart(combinazioni, &combinazione_cumulata, 0, &parts[0], parts, y)

		if parts[0].pointer == len(parts[0].combinazioni) {
			break
		}
	}
}

func gotoNextPart(combinazioni []string, combinazione_cumulata *string, n int, part *part, parts []part, y int) {

	if n+1 < len(parts) { // Move forward until the last part is reached.

		*combinazione_cumulata += string(part.combinazioni[part.pointer])

		n++
		part = &parts[n]
		gotoNextPart(combinazioni, combinazione_cumulata, n, part, parts, y)
	} else { // Move backward: the last part has been reached.

		for _, combinazione := range part.combinazioni {
			fmt.Println(*combinazione_cumulata + string(combinazione))

			tmp := string(*combinazione_cumulata + string(combinazione))
			combinazioni = append(combinazioni, tmp)
		}

		n--
		part = &parts[n]

		if part.pointer+1 < len(part.combinazioni) {
			part.pointer++
		} else {
			part.pointer = 0
			parts[n-1].pointer++
		}
	}
}
