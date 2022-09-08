package main

// https://leetcode.com/problems/missing-number/

import (
	"fmt"
)

func missingNumber(nums []int) int {
	
	sum := 0

	for i:=0; i<=len(nums); i++ {
		sum ^= i
	}

	for _, x := range nums {
		sum ^= x
	}
	
	return sum
}

func main() {

	fmt.Println(missingNumber([]int{3, 0, 1}))
	

	return
}
