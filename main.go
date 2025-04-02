package main

import "fmt"

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 {
		return 0
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
