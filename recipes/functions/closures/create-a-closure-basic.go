// A closure is a special type of anonymous function that
// references variables declared outside the function itself
package main

import "fmt"

func main() {
	// Declaring the variable
	init := 0

	// Assigning an anonymous function to a variable this closure
	counter := func() int {
		init += 1
		return init
	}
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
}

/*

Explanation: The variable "init" was not passed as a parameter to the anonymous function
but the function has access to it. In this example, there is a slight problem as
any other function which will be defined in the main has access to the global variable
"init" and it can change it without calling counter function. Thus, closure also provides
another aspect which is data isolation.
*/
