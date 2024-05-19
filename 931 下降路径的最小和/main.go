package main

func main() {
	input := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	println(minFallingPathSum(input))
}

func minFallingPathSum2(matrix [][]int) int {
	// 初始化dp数组
	n := len(matrix)
	dp := make([][]int, n)
	for row := range dp {
		dp[row] = make([]int, n)
	}
	if n == 1 {
		return matrix[0][0]
	}
	// 动态规划
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = matrix[i][j]
			} else if j == 0 {
				dp[i][j] = minIntSlice([]int{dp[i-1][j], dp[i-1][j+1]}) + matrix[i][j]
			} else if j == n-1 {
				dp[i][j] = minIntSlice([]int{dp[i-1][j-1], dp[i-1][j]}) + matrix[i][j]
			} else {
				dp[i][j] = minIntSlice([]int{dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]}) + matrix[i][j]
			}
		}
	}
	result := minIntSlice(dp[n-1])
	return result

}

func minFallingPathSum(matrix [][]int) int {
	// 初始化dp数组
	n := len(matrix)
	dp := make([][]int, n+1)
	for row := range dp {
		dp[row] = make([]int, n)
	}
	if n == 1 {
		return matrix[0][0]
	}
	// 动态规划
	for i := 1; i < n+1; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[i][j] = minIntSlice([]int{dp[i-1][j] + matrix[i-1][j], dp[i-1][j+1] + matrix[i-1][j+1]}) // todo 如何避免写这么多对应位置上的dp+matrix的语句
			} else if j == n-1 {
				dp[i][j] = minIntSlice([]int{dp[i-1][j-1] + matrix[i-1][j-1], dp[i-1][j] + matrix[i-1][j]})
			} else {
				dp[i][j] = minIntSlice([]int{dp[i-1][j-1] + matrix[i-1][j-1], dp[i-1][j] + matrix[i-1][j], dp[i-1][j+1] + matrix[i-1][j+1]})
			}
		}
	}
	// 第0行全为0可以利用起来，但重写一遍上述逻辑划不来
	result := minIntSlice(dp[n])
	return result

}
func minIntSlice(input []int) int {
	result := input[0]
	for i := range input {
		if result > input[i] {
			result = input[i]
		}
	}
	return result
}
