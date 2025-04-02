package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	counterMap := make(map[int]int, len(nums)+1)

	for _, num := range nums {
		if _, exists := counterMap[num]; exists {
			counterMap[num]++
			continue
		}
		counterMap[num] = 1
	}

	bucket := make([][]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		bucket[i] = make([]int, 0)
	}

	for key, val := range counterMap {
		bucket[val] = append(bucket[val], key)
	}

	result := make([]int, 0)
	for j := len(nums); j > -1 && len(result) < k; j-- {
		if len(bucket[j]) > 0 {
			for _, num := range bucket[j] {
				if len(result) == k {
					return result
				}
				result = append(result, num)
			}
		}
	}

	return result
}

func main() {
	nums := []int{1 /* , 1, 1, 2, 2, 3 */}
	/* k := 2 */
	k := 1

	fmt.Println(topKFrequent(nums, k))
}
