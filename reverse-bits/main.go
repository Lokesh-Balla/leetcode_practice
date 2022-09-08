package main

// https://leetcode.com/problems/reverse-bits/

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
    var ans uint32

	for i := 0; i < 32; i++ {
        ans = ans << 1
		if num & 1 > 0 {
            ans = ans | 1
		}
		num = num >> 1
	}

	return ans
}

func main() {
	fmt.Println(reverseBits(10))

	return
}
