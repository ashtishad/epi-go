package epi

var dialPad = map[rune][]rune{
	'2': []rune{'a', 'b', 'c'},
	'3': []rune{'d', 'e', 'f'},
	'4': []rune{'g', 'h', 'i'},
	'5': []rune{'j', 'k', 'l'},
	'6': []rune{'m', 'n', 'o'},
	'7': []rune{'p', 'q', 'r', 's'},
	'8': []rune{'t', 'u', 'v'},
	'9': []rune{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return recursion([]rune(digits), []rune{})
}

func recursion(digits []rune, s []rune) []string {
	var combos []string

	if len(digits) < 1 {
		return []string{string(s)}
	}

	for _, c := range dialPad[digits[0]] {
		combos = append(combos, recursion(digits[1:], append(s, c))...)
	}
	return combos
}
