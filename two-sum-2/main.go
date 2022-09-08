package main

/*
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted
*/


import "fmt"

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func twoSumIndices(numbers []int, target int, start int, end int) []int {
	if start == end {
		return nil
	}

	// skipping all the positive numbers
	if target < 0 {

		if numbers[end] > 0 {
			return twoSumIndices(numbers, target, start, end-1)
		}

		if abs(numbers[start]+numbers[end]) > abs(target) {
			return twoSumIndices(numbers, target, start, end-1)
		} else if abs(numbers[start]+numbers[end]) < abs(target) {
			return twoSumIndices(numbers, target, start+1, end)
		}

	} else {
		if numbers[start]+numbers[end] > target {
			return twoSumIndices(numbers, target, start, end-1)
		} else if numbers[start]+numbers[end] < target {
			return twoSumIndices(numbers, target, start+1, end)
		}
	}

	return []int{start + 1, end + 1}
}

func twoSum(numbers []int, target int) []int {
	return twoSumIndices(numbers, target, 0, len(numbers)-1)
}

func main() {

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{2, 3, 4}, 7))
	fmt.Println(twoSum([]int{-1, -1, -2, 2, 3, 4}, -2))
}
