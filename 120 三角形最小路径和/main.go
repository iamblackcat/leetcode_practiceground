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
	n := len(triangle[m-1])
	// 初始化dp数组
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}
	// 边缘状态
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
	}
	// 动态规划
	for i := 1; i < len(triangle); i++ {
		for j := 1; j < len(triangle[i])-1; j++ {
			dp[i][j] = minInt(dp[i-1][j]+triangle[i][j], dp[i-1][j-1]+triangle[i][j])
		}
		// 处理三角形斜边
		dp[i][len(triangle[i])-1] = dp[i-1][len(triangle[i-1])-1] + triangle[i][len(triangle[i])-1]
	}
	result := dp[m-1][0]
	for i := 0; i < n; i++ {
		result = minInt(result, dp[m-1][i])
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
