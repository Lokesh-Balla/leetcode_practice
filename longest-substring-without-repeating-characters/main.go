package main

import (
	"fmt"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	m := map[byte]bool{}
	max := 0
	leftP := 0

	for rightP := 0; rightP < len(s); rightP++ {
		character := s[rightP]

		// remove all characters from set till the duplicate is removed
		for m[character] {
			m[s[leftP]] = false
			leftP += 1
		}
		// add the character to the string
		m[character] = true

		// calculate the length and update max
		if max < (rightP - leftP + 1) {
			max = (rightP - leftP + 1)
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
