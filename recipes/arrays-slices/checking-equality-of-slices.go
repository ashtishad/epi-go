package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	// converting string to byte slice for efficiency.
	// strings are immutable in go, ByteSlice aren't
	var str1 string = "userLoginCount" // camelcase
	var str2 string = "UserLoginCount" // PascalCase
	var bs1 = []byte(str1)
	var bs2 = []byte(str2)

	// Optimized equal for byte slices.
	// also treats nil arguments as equivalent to empty slices.
	var equal1 = bytes.Compare(bs1, bs2) == 0
	//equal1 is false

	var equal2 = reflect.DeepEqual(bs1, bs2)
	// recursive comparison. not efficient. but useful for testing
	// applicable for slices,arrays,struct,map,pointer,func,interface
	//equal2 is false

	var equal3 = bytes.Compare(bs1, bs2)
	//equal3 is 1

	//fmt.Println(bs1==bs2) //produce compile-time error
	fmt.Println("checking if nil : ", bs1 == nil) //okay

	fmt.Println(equal1)
	fmt.Println("Deeply equal : ", equal2)
	fmt.Println(equal3)
	var result = Equal(bs1, bs2)
	fmt.Println("using function : ", result)
}

func Equal(string1, string2 []byte) bool {
	if len(string1) != len(string2) {
		return false
	}
	for i, v := range string1 {
		if v != string2[i] {
			return false
		}
	}
	return true
}
