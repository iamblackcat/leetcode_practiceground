package main

func main() {
	println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	// 初始化dp数组
	length := len(s)
	dp := make([][]int, length)
	for row := range dp {
		dp[row] = make([]int, length)
	}
	// 临界状态长度为1或者2时
	for i := 0; i < length; i++ {
		dp[i][i] = 1
	}
	for i := 0; i < length-1; i++ {
		j := i + 1
		if s[i] == s[j] {
			dp[i][j] = 1
		}
	}
	// 动态规划
	for l := 3; l < length; l++ { // length即截取的长度，从短的开始
		for i := 0; i < length-l; i++ { // i为起始位置
			j := i + l - 1
			if s[i] == s[j] && dp[i+1][j-1] == 1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}

		}
	}
	// 得出结果
	result := 0
	resultIndex := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if dp[i][j] == 1 {
				if j-i > result {
					result = j - i
					resultIndex = i
				}
			}
		}
	}

	return s[resultIndex : resultIndex+result+1]
}
