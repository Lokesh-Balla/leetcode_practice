package main

import "fmt"

// https://leetcode.com/problems/longest-consecutive-sequence/

func longestConsecutive(nums []int) int {
	max := 0

	hashSet := map[int]bool{}
	for _, num := range nums {
		hashSet[num] = true
	}

	startPoints := map[int]bool{}
	for k, _ := range hashSet {
		if !hashSet[k-1] {
			startPoints[k] = true
		}
	}

	for k, _ := range startPoints {
		sum := 1
		i := k
		for hashSet[i+1] {
			i++
			sum++
			if sum > max {
				max = sum
			}
		}
		if sum > max {
			max = sum
		}
	}

	return max
}

func main() {

	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{0}))
	return
}
