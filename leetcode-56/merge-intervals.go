package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{
		intervals[0],
	}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > result[len(result)-1][1] {
			result = append(result, intervals[i])
			continue
		}

		if intervals[i][1] >= result[len(result)-1][1] {
			result[len(result)-1][1] = intervals[i][1]
			continue
		}

		if intervals[i][1] < result[len(result)-1][1] {
			continue
		}
	}

	return result
}
