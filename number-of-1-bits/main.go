package main

// https://leetcode.com/problems/number-of-1-bits/

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	weight := 0

	for num != 0 {
		if num & 1 > 0 {
			weight++
		}
		num = num >> 1
	}
	
	return weight
}

func main() {
	fmt.Println(hammingWeight(5))
}
