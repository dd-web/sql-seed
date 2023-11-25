package types

import (
	"fmt"
	"math/rand"
)

var HrSplit string = "\n-----------------------------------------------------\n"
var StdCtrl string = "\033[G\033[K"

// returns a random number between min and max
func RandomBetween(min, max int) int {
	return rand.Intn(max-min) + min
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
