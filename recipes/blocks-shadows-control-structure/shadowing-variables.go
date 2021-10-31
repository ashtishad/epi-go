package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x) // 10 : Shadowed variable
		x := 5
		fmt.Println(x) // 5 : Shadowing variable
	}
	fmt.Println(x) // 10 : Shadowed variable
}

// Learning GO book by john bodner - page 100
// detecting shadowing variable - use "shadow linter" tool
