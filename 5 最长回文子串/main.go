package main

func main() {
	println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	resultL := 0
	resultI := 0

	// 初始化dp数组
	length := len(s)
	dp := make([][]int, length)
	for row := range dp {
		dp[row] = make([]int, length)
	}

	// 动态规划
	for l := 0; l < length; l++ { // length即截取的长度，从短的开始
		for i := 0; i <= length-l-1; i++ { // i为起始位置
			if l == 0 { //
				dp[i][i] = 1
				resultL = 0
				resultI = i
			} else if l == 1 {
				j := i + 1
				if s[i] == s[j] {
					dp[i][j] = 1
					if j-i > resultL {
						resultL = j - i
						resultI = i
					}
				}
			} else {
				j := i + l
				if s[i] == s[j] && dp[i+1][j-1] == 1 {
					dp[i][j] = 1
					if j-i > resultL {
						resultL = j - i
						resultI = i
					}
				} else {
					dp[i][j] = 0
				}
			}

		}
	}

	return s[resultI : resultL+resultI+1]
}
