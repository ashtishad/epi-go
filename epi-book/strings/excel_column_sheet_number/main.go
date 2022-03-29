package main

func titleToNumber(t string) int {
	var res int
	for i := 0; i < len(t); i++ {
		res = res*26 + int(t[i]-'A') + 1
	}
	return res
}
