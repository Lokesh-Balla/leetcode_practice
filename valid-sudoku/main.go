package main

import "fmt"

// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {

	hsr := [9]map[byte]bool{}
	hsc := [9]map[byte]bool{}
	hsb := map[string]map[byte]bool{}
	for i := 0; i < 9; i++ {
		hsr[i] = map[byte]bool{}
		hsc[i] = map[byte]bool{}
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board); c++ {
			if board[r][c] == '.' {
				continue
			}
			if _, ok := hsb[fmt.Sprintf("%v,%v", r/3, c/3)]; !ok {
				hsb[fmt.Sprintf("%v,%v", r/3, c/3)] = map[byte]bool{}
			}

			if hsr[r][board[r][c]] || hsc[c][board[r][c]] || hsb[fmt.Sprintf("%v,%v", r/3, c/3)][board[r][c]] {
				return false
			}

			hsr[r][board[r][c]] = true
			hsc[c][board[r][c]] = true
			hsb[fmt.Sprintf("%v,%v", r/3, c/3)][board[r][c]] = true
		}
	}

	return true
}

func main() {

	return
}
