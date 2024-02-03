package main

import "fmt"

// https://leetcode.com/problems/word-search

var directions = [][2]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func dfs(board [][]byte, x, y, m, n int, index int, word string) bool {
	if index == len(word) {
		return true
	}

	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}

	if word[index] == board[x][y] {

		board[x][y] = '0'

		for _, coordinate := range directions {
			if dfs(board, x+coordinate[0], y+coordinate[1], m, n, index+1, word) {
				return true
			}
		}

		board[x][y] = word[index]
	}

	return false
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 {
		return false
	}

	m := len(board)
	for x, row := range board {
		n := len(row)
		for y := range row {
			if board[x][y] == word[0] {
				if dfs(board, x, y, m, n, 0, word) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	fmt.Println(exist(board, "ABCCED"))

	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCB"))

	board = [][]byte{{'a'}}
	fmt.Println(exist(board, "a"))

	board = [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}
	fmt.Println(exist(board, "AAB"))
}
