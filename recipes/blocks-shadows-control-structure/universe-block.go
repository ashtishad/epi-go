package main

import "fmt"

func main() {
	fmt.Println(true)      // true
	fmt.Printf("%T", true) // bool
	true := 10
	fmt.Println(true) // 10

	fmt.Printf("%T", true) // int
}
