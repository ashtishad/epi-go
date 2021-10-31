package main

import "fmt"

func main() {
	var x int = 10
	var ptrToX *int = &x // & is address operator
	/*
		The * is the indirection operator. It precedes a variable of pointer type and
		returns the pointed-to value. This is called dereferencing.
		Before dereferencing make sure pointer is not nil.
	*/

	if ptrToX != nil {
		fmt.Println(ptrToX)  // prints address -> 0xc0000140c0
		fmt.Println(*ptrToX) // prints pointed-to-value 10
	}

	x = 13
	fmt.Println(x) // 13

	x = *ptrToX + 5
	fmt.Println("X is now : ", x) // 18

	// built-in function new creates a pointer variable. It returns a pointer to a
	// zero value instance of the provided type
	var y = new(int)
	fmt.Println(y == nil) // prints false
	fmt.Println(*y)
	// prints 0
}
