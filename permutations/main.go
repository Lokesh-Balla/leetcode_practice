package main

import "fmt"

// https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	results := [][]int{}

	var dfs func(nums []int, index int)
	dfs = func(nums []int, index int) {
		if index == len(nums) {
			r := make([]int, len(nums))
			copy(r, nums)
			results = append(results, r)
			return
		}

		for i := index; i < len(nums); i++ {
			tmp := nums[i]
			nums[i] = nums[index]
			nums[index] = tmp

			dfs(nums, index+1)

			tmp = nums[i]
			nums[i] = nums[index]
			nums[index] = tmp
		}
	}
	dfs(nums, 0)
	return results
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))

	fmt.Println(permute([]int{1, 2, 3, 4}))
}
