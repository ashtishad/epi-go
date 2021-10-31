package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4}
	var s2 = []int{5, 6, 7}
	var mergedSlices = append(s1, s2...)
	s1 = append(s1, s2...)

	fmt.Println("After merging : ", mergedSlices)
	fmt.Println("s1 becomes : ", s1)
}
