package main

import "fmt"

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}
func update(px *int) {
	*px = 20 // Dereferencing
}
func main() {
	/* CASE 1
	Since the memory location was passed to the function via call-by-value, we
	can’t change the memory address.Thus,You can only reassign the value if there was a
	value already assigned to the pointer
	*/
	var f *int
	failedUpdate(f)
	fmt.Println(f) // prints <nil>
	/*
		CASE 2
		if you want the value assigned to a pointer parameter to still be there when you exit the function,
		you must dereference the pointer and set the value. If you change the pointer,
		you have changed the copy, not the original. Dereferencing puts the new
		value in the memory location pointed to by both the original and the copy.
	*/
	x := 10
	failedUpdate(&x)
	fmt.Println(x) // prints 10
	update(&x)
	fmt.Println(x) // prints 20
}
