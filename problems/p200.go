package problems

// 岛屿的数量 https://leetcode-cn.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}

	return res
}

func dfs(grid [][]byte, i, j int)  {
	rows := len(grid)
	cols := len(grid[0])
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j - 1)
	dfs(grid, i, j + 1)
}