package main

// https://leetcode.com/problems/sum-of-two-integers/

import (
	"fmt"
)

func getSum(a int, b int) int {
	if b == 0 {
		return a
	}

	return getSum(a^b, (a&b)<<1)
}

func main() {

	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(2, 3))
	fmt.Println(getSum(12, 3))
	fmt.Println(getSum(-1, 1))

	return
}
