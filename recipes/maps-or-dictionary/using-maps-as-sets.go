package main

import "fmt"

func main() {
	intSet := map[int]bool{}
	inputs := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range inputs {
		intSet[v] = true
	}
	fmt.Println(len(inputs), len(intSet)) // value = 11 , set = 8
	fmt.Println("Full set : ", intSet)
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}
