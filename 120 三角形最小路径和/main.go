package main

func main() {

	input := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	println(minimumTotal(input))

}

func minimumTotal(triangle [][]int) int {
	m := len(triangle)

	// 初始化dp数组
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 0
	}
	// 动态规划
	dp[0] = triangle[0][0]
	for i := 1; i < m; i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {

			if j == 0 {
				// 左侧边时
				dp[0] += triangle[i][0]
			} else if j == len(triangle[i])-1 {
				// 右侧边时
				dp[j] = triangle[i][j] + dp[j-1]
			} else {
				dp[j] = minInt(dp[j], dp[j-1]) + triangle[i][j]

			}
		}
	}
	result := dp[0]
	for i := 1; i < m; i++ {
		result = minInt(result, dp[i])
	}
	return result
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
