package main

import "fmt"

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 {
		return 0
	}

	for i := 1; i < len(heightMap)-1; i++ {
		for j := 1; j < len(heightMap[i])-1; j++ {
			fmt.Println(i, j)
		}
	}

	return 0
}

func main() {
	inputs := [][][]int{
		{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}},
	}

	for _, input := range inputs {
		fmt.Println(trapRainWater(input))
	}
}
