package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
https://leetcode.com/problems/valid-palindrome/
*/

func checkPalindrome(s []rune) bool {
	if len(s) <= 1 {
		return true
	}
	
	if s[0] != s[len(s)-1] {
		return false
	}

	return checkPalindrome(s[1 : len(s)-1])
}

func isPalindrome(s string) bool {

	s = strings.ToLower(s)

	r := []rune{}

	for _, x := range s {
		if unicode.IsDigit(x) || unicode.IsLetter(x) {
			r = append(r, x)
		}
	}

	return checkPalindrome(r)
}

func main() {

	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	return
}
