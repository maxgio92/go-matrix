package main

import "fmt"

type combinazione string
type part struct {
	ordinata     int
	combinazioni []combinazione
}

var (
	parts = []part{
		part{combinazioni: []combinazione{"A", "B"}, ordinata: 0},                // part
		part{combinazioni: []combinazione{"1", "2", "3", "4", "5"}, ordinata: 0}, // part
		part{combinazioni: []combinazione{"w", "x", "y", "z"}, ordinata: 0},      // part
		part{combinazioni: []combinazione{"E", "F", "G", "H"}, ordinata: 0},      // part
		part{combinazioni: []combinazione{"A", "B"}, ordinata: 0},                // part
	}
	endMarker = false
)

func main() {
	for _, v := range parts {
		fmt.Println(v.combinazioni)
	}
	combinazioni := getCombinazioni()
	fmt.Println(combinazioni)
}

// (ordinata)
// y
// ^	A	1	Z
// |	B	2	Y
// |		3	X
// ------> x (ascissa)

func getCombinazioni() []string {
	combinazioni := []string{}
	combinazione_cumulata := ""

	// For each time the last part has been reached
	// exit from recursion until reaching this:
	for {
		combinazione_cumulata = ""

		// Start always from the first part (x=0)
		gotoNextPart(&combinazioni, &combinazione_cumulata, 0, &parts[0], parts)

		if parts[0].ordinata == len(parts[0].combinazioni) || endMarker {
			break
		}
	}

	return combinazioni
}

func gotoNextPart(combinazioni *[]string, combinazione_cumulata *string, ascissa int, part *part, parts []part) {

	if ascissa+1 < len(parts) { // Until the last part is reached

		*combinazione_cumulata += string(part.combinazioni[part.ordinata])

		// Move forward
		ascissa++
		part = &parts[ascissa]
		gotoNextPart(combinazioni, combinazione_cumulata, ascissa, part, parts)
	} else { // When the last part is reached

		for _, combinazione := range part.combinazioni {
			*combinazioni = append(*combinazioni, string(*combinazione_cumulata+string(combinazione)))
		}

		// Move backward
		ascissa--
		part = &parts[ascissa]

		// Store where we gone
		scorriCombinazionePartePrecedente(part, parts, ascissa)
	}
}

func scorriCombinazionePartePrecedente(part *part, parts []part, ascissa int) {
	if part.ordinata+1 < len(part.combinazioni) {
		part.ordinata++
	} else {
		part.ordinata = 0
		ascissa--
		if ascissa >= 0 {
			scorriCombinazionePartePrecedente(&parts[ascissa], parts, ascissa)
		} else {
			endMarker = true
		}
	}
}
