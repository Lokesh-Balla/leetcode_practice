package main

import "fmt"

// https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	var dfs func(index int, sum int, arr []int)
	dfs = func(index int, sum int, arr []int) {
		if sum == target {
			result = append(result, arr)
			return
		}

		if index >= len(candidates) || sum > target {
			return
		}

		dfs(index, sum+candidates[index], append([]int{candidates[index]}, arr...))

		dfs(index+1, sum, arr)
	}

	dfs(0, 0, []int{})

	return result
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
