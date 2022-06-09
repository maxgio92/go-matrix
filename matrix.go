package main

import "fmt"

type combination string
type part struct {
	ordinate     int
	combinations []combination
}

var (
	parts = []part{
		part{combinations: []combination{"A", "B"}, ordinate: 0},                // part
		part{combinations: []combination{"1", "2", "3", "4", "5"}, ordinate: 0}, // part
		part{combinations: []combination{"w", "x", "y", "z"}, ordinate: 0},      // part
		part{combinations: []combination{"E", "F", "G", "H"}, ordinate: 0},      // part
		part{combinations: []combination{"A", "B"}, ordinate: 0},                // part
	}
	endMarker = false
)

// (ordinate)
// y
// ^	A	1	Z
// |	B	2	Y
// |		3	X
// ------> x (abscissa)

func main() {
	for _, v := range parts {
		fmt.Println(v.combinations)
	}
	combinations := getcombinations()
	fmt.Println(combinations)
}

func getcombinations() []string {
	combinations := []string{}
	combination_sum := ""

	// For each time the last part has been reached
	// exit from recursion until reaching this:
	for {
		combination_sum = ""

		// Start always from the first part (x=0)
		gotoNextPart(&combinations, &combination_sum, 0, &parts[0], parts)

		if parts[0].ordinate == len(parts[0].combinations) || endMarker {
			break
		}
	}

	return combinations
}

func gotoNextPart(combinations *[]string, combination_sum *string, abscissa int, part *part, parts []part) {

	if abscissa+1 < len(parts) { // Until the last part is reached

		*combination_sum += string(part.combinations[part.ordinate])

		// Move forward
		abscissa++
		part = &parts[abscissa]
		gotoNextPart(combinations, combination_sum, abscissa, part, parts)

	} else { // When the last part is reached

		for _, combination := range part.combinations {
			*combinations = append(*combinations, string(*combination_sum+string(combination)))
		}

		// Move backward
		abscissa--
		part = &parts[abscissa]

		// Store where we gone
		scrollDownPrevPartCombination(part, parts, abscissa)
	}
}

func scrollDownPrevPartCombination(part *part, parts []part, abscissa int) {

	if part.ordinate+1 < len(part.combinations) {
		part.ordinate++

	} else {
		part.ordinate = 0
		abscissa--

		if abscissa >= 0 {
			scrollDownPrevPartCombination(&parts[abscissa], parts, abscissa)
		} else {
			endMarker = true
		}
	}
}
