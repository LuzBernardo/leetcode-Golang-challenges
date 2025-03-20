package main

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i > len(grid[0])-1 || j > len(grid)-1 || grid[j][i] != '1' {
		return
	}

	grid[j][i] = '0'
	dfs(grid, i-1, j) // esquerdo
	dfs(grid, i, j-1) // embaixo
	dfs(grid, i+1, j) // direito
	dfs(grid, i, j+1) // em cima
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	counter := 0

	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[j]); i++ {
			if grid[j][i] == '1' {
				counter++
				dfs(grid, i, j)
			}
		}
	}

	return counter
}
