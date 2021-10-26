package main

func Median(a, b, c int) int {
	var med = max(min(a, b), min(max(a, b), c))
	return med
}
