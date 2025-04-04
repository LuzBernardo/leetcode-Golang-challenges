package main

import (
	"fmt"
)

type height struct {
	height int
	i      int
	j      int
}

type heap struct {
	heights []*height
}

func (h *heap) Len() int {
	return len(h.heights)
}

func (h *heap) heapifyUp(index int) {

	for index > 0 {
		parent := (index - 1) / 2
		if h.heights[parent].height > h.heights[index].height {
			h.heights[parent], h.heights[index] = h.heights[index], h.heights[parent]
			index = parent
		} else {
			return
		}
	}
}

func (h *heap) heapifyDown(index int) {
	for {
		left := index*2 + 1
		right := index*2 + 2
		smallest := index
		if left < len(h.heights) && h.heights[left].height < h.heights[smallest].height {
			smallest = left
		}
		if right < len(h.heights) && h.heights[right].height < h.heights[smallest].height {
			smallest = right
		}
		if smallest == index {
			break
		}

		h.heights[smallest], h.heights[index] = h.heights[index], h.heights[smallest]
		index = smallest
	}
}

func (h *heap) Pop() *height {
	if len(h.heights) == 0 {
		return nil
	}

	first := h.heights[0]

	h.heights[0] = h.heights[len(h.heights)-1]
	h.heights = h.heights[:len(h.heights)-1]

	h.heapifyDown(0)

	return first
}

func (h *heap) Push(cell height) {
	h.heights = append(h.heights, &cell)
	h.heapifyUp(len(h.heights) - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}
	visited := make([][]bool, len(heightMap))
	water := 0

	for i := 0; i < len(heightMap); i++ {
		visited[i] = make([]bool, len(heightMap[i]))
	}

	h := heap{
		heights: make([]*height, 0),
	}

	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for j := 0; j < len(heightMap[0]); j++ {
		h.Push(height{heightMap[0][j], 0, j})
		h.Push(height{heightMap[len(heightMap)-1][j], len(heightMap) - 1, j})

		visited[0][j] = true
		visited[len(heightMap)-1][j] = true
	}
	for i := 1; i < len(heightMap); i++ {
		first, last := 0, len(heightMap[i])-1

		h.Push(height{heightMap[i][first], i, first})
		h.Push(height{heightMap[i][last], i, last})

		visited[i][first] = true
		visited[i][last] = true
	}

	for h.Len() > 0 {
		smallest := h.Pop()
		i, j := smallest.i, smallest.j
		currHeight := smallest.height

		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]

			if ni < len(heightMap) && ni >= 0 && nj < len(heightMap[0]) &&
				nj >= 0 && !visited[ni][nj] {

				visited[ni][nj] = true
				nextHeight := heightMap[ni][nj]

				delta := currHeight - nextHeight
				if delta > 0 {
					water += delta
				}

				h.Push(height{max(currHeight, nextHeight), ni, nj})
			}
		}
	}

	return water
}

func main() {
	inputs := [][][]int{
		{
			{1, 4, 3, 1, 3, 2},
			{3, 2, 1, 3, 2, 4},
			{2, 3, 3, 2, 3, 1},
		},
		{
			{1, 4, 3, 1},
			{3, 2, 1, 3},
			{2, 3, 1, 2},
			{1, 3, 2, 1}},
	}

	for _, input := range inputs {
		fmt.Println(trapRainWater(input))
	}
}
