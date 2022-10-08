package main

func getMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxArea(height []int) int {

	li := 0
	ri := len(height) - 1
	max := -1

	for li <= ri {
		area := getMin(height[li], height[ri]) * (ri - li)
		if area > max {
			max = area
		}

		if height[li] < height[ri] {
			li++
		} else {
			ri--
		}
	}

	return max
}