package main

func main() {
	input := [][]int{
		{0, 1},
		{0, 0},
	}

	println(uniquePathsWithObstacles(input))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var dp [][]int
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	// 初始化dp
	dp = make([][]int, m)
	for row := range obstacleGrid {
		dp[row] = make([]int, n)
	}
	// 起始或终点被石头堵上
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	} else {
		dp[0][0] = 1
	}
	// 动态规划
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if obstacleGrid[row][col] == 1 {
				dp[row][col] = 0
				continue
			}
			// 网格起始状态
			if row == 0 && col == 0 {
				continue
			}

			// 边缘状态
			if row == 0 {
				if dp[row][col-1] != 0 {
					dp[row][col] = 1
				} else {
					dp[row][col] = 0
				}
			} else if col == 0 {
				if dp[row-1][col] != 0 {
					dp[row][col] = 1
				} else {
					dp[row][col] = 0
				}
			} else {
				dp[row][col] = dp[row-1][col] + dp[row][col-1]
			}

		}
	}
	return dp[m-1][n-1]
}
