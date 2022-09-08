package main

// https://leetcode.com/problems/single-number/

import (
	"fmt"
)

func singleNumber(nums []int) int {
   	if len(nums) == 1 {
		return nums[0]
	}

	sumXOR := nums[0]

	for i := 1; i < len(nums); i++ {
		sumXOR = sumXOR ^ nums[i]
	}

	return sumXOR
}

func main() {
	fmt.Println(singleNumber([]int{2,2,1}))
	fmt.Println(singleNumber([]int{4,1,2,1,2})) 
	fmt.Println(singleNumber([]int{1}))
}

