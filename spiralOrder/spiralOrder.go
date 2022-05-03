package spiralOrder

func spiralOrder(matrix [][]int) []int {
	h := make(map[int]map[int]bool)
	var res []int
	m := len(matrix)
	n := len(matrix[0])
	i, j := 0, 0
	var forward int
	for index := range matrix {
		h[index] = make(map[int]bool)
	}
	for {
		res = append(res, matrix[i][j])
		h[i][j] = true
		if forward == 0 {
			if j+1 < n && !h[i][j+1] {
				j++
				continue
			}
			if (j+1 >= n || h[i][j+1]) && (i+1 < m && !h[i+1][j]) {
				forward = 1
				i++
				continue
			}
			break
		}
		if forward == 1 {
			if i+1 < m && !h[i+1][j] {
				i++
				continue
			}
			if (i+1 >= m || h[i+1][j]) && (j-1 >= 0 && !h[i][j-1]) {
				forward = 2
				j--
				continue
			}
			break
		}
		if forward == 2 {
			if j-1 >= 0 && !h[i][j-1] {
				j--
				continue
			}
			if (j-1 < 0 || h[i][j]) && (i-1 >= 0 && !h[i-1][j]) {
				forward = 3
				i--
				continue
			}
			break
		}
		if forward == 3 {
			if i-1 >= 0 && !h[i-1][j] {
				i--
				continue
			}
			if (i-1 < 0 || h[i-1][j]) && (j+1 < n && !h[i][j+1]) {
				forward = 0
				j++
				continue
			}
			break
		}
	}
	return res
}
