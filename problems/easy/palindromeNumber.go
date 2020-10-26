package problems

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	chars := []rune(str)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	s1 := string(chars)
	if str == s1 {
		return true
	}

	return false
}
