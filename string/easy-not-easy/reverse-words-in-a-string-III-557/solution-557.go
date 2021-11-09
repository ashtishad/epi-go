package leetcode

import "strings"

func reverseWords(s string) string {
	var res = []string{}
	var words = strings.Split(s, " ")

	for _, word := range words {
		res = append(res, reverse(word))
	}
	return strings.Join(res, " ")
}

func reverse(s string) string {
	var temp = []byte(s)
	var lp, rp = 0, len(temp) - 1
	for lp < rp {
		temp[lp], temp[rp] = temp[rp], temp[lp]
		lp++
		rp--
	}
	return string(temp)
}
