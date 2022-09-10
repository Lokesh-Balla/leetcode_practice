package main

import (
	"fmt"
)

// https://leetcode.com/problems/longest-common-subsequence/

// Solving using recursion and memoisation
// ===============================================================================================
func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func lcs(text1 string, text2 string, index1 int, index2 int, m map[string]int) int {

	key := fmt.Sprintf("%d:%d", index1, index2)

	if val, ok := m[key]; ok {
		return val
	}

	if index1 == len(text1) || index2 == len(text2) {
		return 0
	}

	val := 0
	if text1[index1] == text2[index2] {
		val = 1 + lcs(text1, text2, index1+1, index2+1, m)

	} else {
		val = max(lcs(text1, text2, index1+1, index2, m), lcs(text1, text2, index1, index2+1, m))
	}
	m[key] = val
	return val
}

// ===============================================================================================

// solving using dynamic programming using 2-d matrix
// ===============================================================================================

func init2DSlice(rows, cols int) [][]int {
	mat := make([][]int, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]int, cols)
	}
	return mat
}

func lcsUsing2DMat(text1 string, text2 string) int {
	/*
		The rows and cols will correspond to the charecters from the strings
			a	b	c	d	e
		a [ 3	2	2	1	1 ] 0
		c [	2	2	2	1	1 ] 0
		e [	1	1	1	1	1 ] 0
			0	0	0	0	0	0
		we start caluclating the values from the bottom right corner and process left

		each value is caluclated by first check if the values match
		if match
			1 + value of the mat[x+1][y+1]
		else
			max(mat[x+1][y], mat[x][y+1])

	*/

	maxVal := 0
	rows, cols := len(text1), len(text2)
	mat := init2DSlice(rows+1, cols+1)

	for rIndex := rows - 1; rIndex >= 0; rIndex-- {
		for cIndex := cols - 1; cIndex >= 0; cIndex-- {
			if text1[rIndex] == text2[cIndex] {
				mat[rIndex][cIndex] = 1 + mat[rIndex+1][cIndex+1]
			} else {
				mat[rIndex][cIndex] = max(mat[rIndex+1][cIndex], mat[rIndex][cIndex+1])
			}
			maxVal = max(maxVal, mat[rIndex][cIndex])
		}
	}

	return maxVal
}

// ===============================================================================================

func longestCommonSubsequence(text1 string, text2 string) int {
	return lcsUsing2DMat(text1, text2)

	// return lcs(text1, text2, 0, 0, map[string]int{})
}

func main() {

	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
	fmt.Println(longestCommonSubsequence("bsbininm", "jmjkbkjkv"))

}
