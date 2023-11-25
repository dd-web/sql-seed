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

func GenerateAccount() *Account {
	return &Account{}
}
