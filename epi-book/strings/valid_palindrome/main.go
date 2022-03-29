package epi

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isAlphaNumeric(s[i]) {
			i++
			continue
		}
		if !isAlphaNumeric(s[j]) {
			j--
			continue
		}
		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

// isAlphaNumeric returns true if the character is a letter or digit.
// similar to unicode.IsLetter + unicode.IsNumber
func isAlphaNumeric(b byte) bool {
	if b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return 'a' + b - 'A'
	}
	return b
}
