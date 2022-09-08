package main

// https://leetcode.com/problems/counting-bits/

import (
	"fmt"
)

func countSetBits(n int) int {
	setBits := 0
	for n != 0 {
		if n & 1 > 0 {
			setBits++
		}
		n = n >> 1
	}
	return setBits
}

func countBits(n int) []int {
	ans := []int{}
   	for i := 0; i <= n; i++ {
		ans = append(ans, countSetBits(i))
	} 
	return ans
}

func main() {
	fmt.Println(countBits(5))
}
