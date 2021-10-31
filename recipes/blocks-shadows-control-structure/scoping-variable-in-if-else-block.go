package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if n := rand.Intn(15); n == 0 { // n's scope is only in if block
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	//fmt.Println(n) // out of scope : unidentified
}
