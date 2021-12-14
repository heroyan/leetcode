package problems

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	rows := len(grid1)
	cols := len(grid1[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				dfs4(grid2, i, j, rows, cols)
			}
		}
	}

	total := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid2[i][j] == 1 {
				total++
				dfs4(grid2, i, j, rows, cols)
			}
		}
	}

	return total
}

func dfs4(grid [][]int, i, j, rows, cols int)  {
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs4(grid, i-1, j, rows, cols)
	dfs4(grid, i+1, j, rows, cols)
	dfs4(grid, i, j - 1, rows, cols)
	dfs4(grid, i, j + 1, rows, cols)
}