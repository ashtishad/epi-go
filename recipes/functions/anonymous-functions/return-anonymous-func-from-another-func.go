package main

import "fmt"

// GForG : Returning anonymous function
func GForG() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "GeeksforGeeks"
	}
	return myf
}

func main() {
	value := GForG()
	fmt.Println(value("Welcome ", "to "))
}
