package types

import (
	"fmt"
	"math/rand"
)

var HrSplit string = "\n-----------------------------------------------------\n"
var StdCtrl string = "\033[G\033[K"

// returns a random number between min and max
func RandomBetween[T int | rune](min, max T) T {
	return T(T(rand.Intn(int(max)-int(min))) + T(min))
}

// prints the given string separated by a horizontal line on top and bottom
func HrPrint(str string) {
	fmt.Printf("\n" + HrSplit)
	fmt.Print("  ", str)
	fmt.Printf(HrSplit + "\n")
}

// returns a random weighted item from the given map
// the map should be of the form map[T]int where T is the type of item and int is the weight
// if for some reason a random item cannot be chosen, the first item in the map is returned
// this is why the map should never be empty
func RandomWeightedFromMap[T comparable](weights map[T]int) T {
	var cumulativeWeights []int
	var list []T
	cumulative := 0

	for item, weight := range weights {
		cumulative += weight
		cumulativeWeights = append(cumulativeWeights, cumulative)
		list = append(list, item)
	}

	r := rand.Intn(cumulativeWeights[len(cumulativeWeights)-1])

	for i, weight := range cumulativeWeights {
		if r < weight {
			return list[i]
		}
	}

	return list[0]
}
