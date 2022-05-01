package main

import "fmt"

func findCircleNum(isConnected [][]int) int {
	var dp = make([][]int, len(isConnected))
	for i := range isConnected {
		dp[i] = append(dp[i], isConnected[i]...)
	}
	// 1、步长
	for i := 2; i < len(dp); i++ {
		// 2、开始的下标
		for j := 0; j < len(dp); j++ {
			// 3、结束的下标
			for k := j + 1; k < j+i && k < len(dp[j]) && j+i < len(dp[j]); k++ {
				if dp[j][k] == 1 && dp[k][j+i] == 1 {
					dp[j][j+i] = 1
					break
				}
			}
		}
	}
	var i int
	var cityNum int = 1
	m := make(map[int]int)
	city := make(map[int]struct{})
	for {
		if i >= len(dp) {
			break
		}
		if _, ok := city[m[i]]; ok {
			continue
		}
		city[cityNum] = struct{}{}
		for node := range dp[i] {
			m[node] = cityNum
		}
	}
	return len(city)
}

func main() {
	fmt.Println(findCircleNum([][]int{{1}, {1}, {1}}))
}
