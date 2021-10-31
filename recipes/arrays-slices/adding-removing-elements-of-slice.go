package main

import "fmt"

func main() {
	var primeNumbers = []int{2, 5, 7}
	primeNumbers = append(primeNumbers, 11)
	//primeNumbers is {2, 5, 7, 11}
	fmt.Println(primeNumbers)

	primeNumbers = append(primeNumbers[0:1],
		append([]int{3}, primeNumbers[1:]...)...)
	//primeNumbers is {2, 3, 5, 7, 11}
	fmt.Println(primeNumbers)

	primeNumbers = append(primeNumbers[0:2], primeNumbers[3:]...)
	//primeNumbers is {2, 3, 7, 11}
	fmt.Println(primeNumbers)

	primeNumbers = append(primeNumbers, 13, 17)
	//primeNumbers is {2, 3, 7, 11, 13, 17}
	fmt.Println(primeNumbers)

	primeNumbers = append([]int{2, 3, 5}, primeNumbers[2:]...)
	//primeNumbers is {2, 3, 5, 7, 11, 13, 17}
	fmt.Println(primeNumbers)

	primeNumbers = primeNumbers[:0]
	//primeNumbers is {}
	fmt.Println(primeNumbers)
}
