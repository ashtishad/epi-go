package main

import "fmt"

func main() {
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName: "Pat",
		//MiddleName: "Perry", // This line won't compile
		MiddleName: stringToPointer("Perry"), // This works
		LastName:   "Peterson",
	}

	fmt.Println(p.MiddleName)  // 0xc000010230
	fmt.Println(*p.MiddleName) // Perry
}

// tip: Use a helper function to turn a constant value into a pointer.
// returns pointed address
func stringToPointer(s string) *string {
	return &s
}

/*
Why does stringToPointer function work?
When we pass a constant to a function, the constant is
copied to a parameter, which is a variable. Since it’s a variable, it has an
address in memory. The function then returns the variable’s memory address
*/
