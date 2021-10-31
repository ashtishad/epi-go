package main

import "fmt"

func main() {
	//creating
	var s []int // best practice: when initializing, always assign type
	s1 := []string{"foo", "bar"}
	s2 := make([]int, 2)    // same as []int{0, 0}
	s3 := make([]int, 2, 4) // same as new([4]int)[:2]
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3)) // 2 4

	// slicing
	a := [...]int{0, 1, 2, 3} // an array
	s = a[1:3]                // s == []int{1, 2}        cap(s) == 3
	s = a[:2]                 // s == []int{0, 1}        cap(s) == 4
	s = a[2:]                 // s == []int{2, 3}        cap(s) == 2
	s = a[:]                  // s == []int{0, 1, 2, 3}  cap(s) == 4

	fmt.Println(s)

	// iterating

	for i, v := range s {
		fmt.Println(i, v)
	}
}
