package main

import "fmt"

func main() {
	/* using len and cap function
	append increases len by 1, if capacity is not full
	if cap = full , it increases cap by 2x(before 1024), after 1024 it increases by 25%
	*/
	var init = []int{} // preferable
	var xl = []int{1, 2, 3}
	y := []int{20, 30, 40}

	fmt.Println(len(init))
	//xl = append(xl, 5, 6, 7)
	xl = append(xl, y...)
	fmt.Println(xl)
	fmt.Println(len(xl), cap(xl))

	/*
		Go is a call by value language. Every time you pass a parameter to a function,
		Go makes a copy of the value that’s passed in. Passing a slice to the
		append function actually passes a copy of the slice to the function.
		The function adds the values to the copy of the slice and returns the copy.
		You then assign the returned slice back to the variable in the calling function.
	*/

	// using make function
	x := make([]int, 5) // prints {0,0,0,0,0} len=5=cap
	x = append(x, 10)   // {0,0,0,0,0, 10} increases len by 1,
	// 10 placed at the end. len=6 cap =10
	fmt.Println(len(x), cap(x))
	//var xll = make([]int, 0, 10) // len=0, cap=10
	//fmt.Println(len(xll), cap(xll))
}
