package main

import "fmt"

// https://leetcode.com/problems/robot-bounded-in-circle/

func isRobotBounded(instructions string) bool {

	if len(instructions) == 0 {
		return true
	}

	direction := []int{0, 1}
	presentCordinaties := []int{0, 0}

	for _, instruction := range instructions {
		if instruction == 'G' {
			presentCordinaties[0] += direction[0]
			presentCordinaties[1] += direction[1]
		} else if instruction == 'L' {
			direction[0], direction[1] = -direction[1], direction[0]
		} else if instruction == 'R' {
			direction[0], direction[1] = direction[1], -direction[0]
		}
	}

	// either end up at 0,0 again or if direction changed then we'll end up at 0,0 eventually
    // in 0, 2, 4 cycles
	return (presentCordinaties[0] == 0 && presentCordinaties[1] == 0) || (direction[0] != 0 || direction[1] != 1)
}

func main() {

	fmt.Println(isRobotBounded("GGLLGG"))
	fmt.Println(isRobotBounded("GG"))
	fmt.Println(isRobotBounded("GL"))
	fmt.Println(isRobotBounded("RL"))

}
