package problems

// 岛屿的数量 https://leetcode-cn.com/problems/number-of-islands/
func closedIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	res := 0
	for i := 0; i < cols; i++ {
		dfs2(grid, 0, i)
		dfs2(grid, rows-1, i)
	}
	for i := 0; i < rows; i++ {
		dfs2(grid, i, 0)
		dfs2(grid, i, cols-1)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				res++
				dfs2(grid, i, j)
			}
		}
	}

	return res
}

func dfs2(grid [][]int, i, j int)  {
	rows := len(grid)
	cols := len(grid[0])
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	dfs2(grid, i-1, j)
	dfs2(grid, i+1, j)
	dfs2(grid, i, j - 1)
	dfs2(grid, i, j + 1)
}

var maxArea int
func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	total := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				maxArea = 0
				dfs3(grid, i, j)
				if total < maxArea {
					total = maxArea
				}
			}
		}
	}

	return total
}

func dfs3(grid [][]int, i, j int)  {
	rows := len(grid)
	cols := len(grid[0])
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	maxArea++
	dfs3(grid, i-1, j)
	dfs3(grid, i+1, j)
	dfs3(grid, i, j - 1)
	dfs3(grid, i, j + 1)
}