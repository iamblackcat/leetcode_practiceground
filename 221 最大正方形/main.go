package main

func main() {
	input := [][]byte{
		{'0'},
	}
	println(maximalSquare(input))
}

func maximalSquare(matrix [][]byte) int {
	// 初始化dp
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, len(matrix))
	for row := range dp {
		dp[row] = make([]int, len(matrix[row]))
	}
	// 特殊情况之0行
	for i := 0; i < n; i++ {
		if matrix[0][i] == '0' {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
	}
	// 特殊情况之0列
	for i := 0; i < m; i++ {
		if matrix[i][0] == '0' {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
	}

	// 动态规划
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = minIntSlice([]int{dp[i-1][j-1], dp[i-1][j], dp[i][j-1]}) + 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	// 寻找最大边长
	maxSeq := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] > maxSeq {
				maxSeq = dp[i][j]
			}
		}
	}

	return maxSeq * maxSeq
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
