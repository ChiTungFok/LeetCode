package numIslands

func numIslands(grid [][]byte) int {
	var islands int
	var path [][]bool
	for i := range grid {
		path = append(path, make([]bool, len(grid[i])))
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !path[i][j] {
				dfs(grid, i, j, path)
				islands++
			}
		}
	}
	return islands
}

func dfs(grid [][]byte, i, j int, path [][]bool) {
	if i >= len(grid) || i < 0 {
		return
	}
	if j >= len(grid[i]) || j < 0 {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	if path[i][j] {
		return
	}
	path[i][j] = true
	dfs(grid, i, j+1, path)
	dfs(grid, i, j-1, path)
	dfs(grid, i+1, j, path)
	dfs(grid, i-1, j, path)
}
