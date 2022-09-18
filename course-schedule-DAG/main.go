package main

import "fmt"

// https://leetcode.com/problems/course-schedule/

func checkLoop(key int, adjacencyList map[int][]int, visited map[int]bool) bool {
	if visited[key] {
		return true
	}

	if val, ok := adjacencyList[key]; ok && len(val) == 0 {
		return false
	}
	
	visited[key] = true
	for _, x := range adjacencyList[key] {
		if checkLoop(x, adjacencyList, visited) {
			return true
		}
	}

	visited[key] = false
	adjacencyList[key] = []int{}
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	// creating a adjacency list
	adjacencyList := map[int][]int{}
	for _, edge := range prerequisites {
		if val, ok := adjacencyList[edge[0]]; ok {
			adjacencyList[edge[0]] = append(val, edge[1])
		} else {
			adjacencyList[edge[0]] = []int{edge[1]}
		}
	}

	for i := 0; i < numCourses; i++ {
		if checkLoop(i, adjacencyList, map[int]bool{}) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
}
