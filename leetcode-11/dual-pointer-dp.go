package main

func areaCalculation(h int, l int) int {
	return h * l
}

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	maxArea, left := 0, 0
	right := len(height) - 1

	for i := 0; i < len(height)-1; i++ {
		if right == left {
			return maxArea
		}
		area := 0

		if height[left] > height[right] {
			area = areaCalculation(height[right], right-left)
			right--
		} else {
			area = areaCalculation(height[left], right-left)
			left++
		}

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
