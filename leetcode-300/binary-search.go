package main

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func lengthOfLIS(nums []int) int {
	tails := make([]int, 0, len(nums))

	for _, num := range nums {
		idx := binarySearch(tails, num)
		if idx == len(tails) {
			tails = append(tails, num)
		} else {
			tails[idx] = num
		}
	}

	return len(tails)
}
