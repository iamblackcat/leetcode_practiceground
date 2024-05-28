package main

func main() {
	println(minDistance("", "a"))
}

func minDistance(word1 string, word2 string) int {
	// 初始化dp数组
	length1 := len(word1)
	length2 := len(word2)
	dp := make([][]int, length1+1)
	for row := range dp {
		dp[row] = make([]int, length2+1)
	}
	for i := 0; i < length2+1; i++ {
		dp[0][i] = i
	}
	for i := 0; i < length1+1; i++ {
		dp[i][0] = i
	}
	// 动态规划遍历
	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j-1]-1, dp[i-1][j]) + 1
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j-1], dp[i-1][j]) + 1
			}

		}
	}

	return dp[length1][length2]
}
