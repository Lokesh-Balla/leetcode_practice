package main

import "fmt"

// https://leetcode.com/problems/subsets

func dfs(s []int, i int, nums []int) [][]int {
	if i == len(nums) {
		return [][]int{s}
	}

	ans := [][]int{}

	sCopy := make([]int, len(s))
	copy(sCopy, s)	

	sCopy = append(sCopy, nums[i])
	ans = append(ans, dfs(sCopy, i+1, nums)...)

	ans = append(ans, dfs(s, i+1, nums)...)

	return ans
}

func subsets(nums []int) [][]int {
	return dfs([]int{}, 0, nums)
}

func main() {

	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))

}
