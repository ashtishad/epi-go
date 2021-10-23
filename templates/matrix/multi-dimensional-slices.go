package main

import "fmt"

func main() {

	// Creating multi-dimensional slice
	s1 := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 1 : ", s1)

	// Creating multi-dimensional slice
	s2 := [][]string{
		[]string{"str1", "str2"},
		[]string{"str3", "str4"},
		[]string{"str5", "str6"},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 2 : ", s2)

	// Printing every slice inside s2
	fmt.Println("MultiDimensional Slice s2:")
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

	// Printing elements in slice as matrix
	fmt.Println("Slice s2 Like Matrix:")

	// number of rows in s2
	n := len(s2)

	// number of columns in s2
	m := len(s2[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(s2[i][j] + " ")
		}
		fmt.Print("\n")
	}
}
