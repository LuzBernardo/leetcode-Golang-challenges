package main

func max(arr []int) int {
	maxHeight := 0
	idx := 0

	for i, height := range arr {
		if height > maxHeight {
			maxHeight = height
			idx = i
		}
	}

	return idx
}

func trap(height []int) int {
	start := 0
	water := 0

	idxHeight := max(height)

	for i := 0; i <= idxHeight; i++ {
		if i == 0 {
			continue
		}
		if height[i] >= height[start] {
			if start != 0 || (start == 0 && height[start] != 0) {
				length := i - start - 1
				litters := length * height[start]
				for j := start + 1; j < i; j++ {
					litters -= height[j]
				}
				water += litters
			}

			start = i
		}
	}

	start = len(height) - 1
	for j := len(height) - 1; j >= idxHeight; j-- {
		if j == len(height)-1 {
			continue
		}
		if height[j] >= height[start] {
			if start != len(height)-1 || (start == len(height)-1 && height[start] > 0) {
				length := start - j - 1
				litters := length * height[start]

				for l := start - 1; l > j; l-- {
					litters -= height[l]
				}

				water += litters
			}

			start = j
		}
	}

	return water
}

func trap2(height []int) int {
	if len(height) < 3 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] <= height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	return water
}
