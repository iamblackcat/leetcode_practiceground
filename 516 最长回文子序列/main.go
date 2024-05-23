package main

func main() {
	println(longestPalindromeSubseq("bbbab"))
}

func longestPalindromeSubseq(s string) int {
	length := len(s)
	// 初始化dp数组
	dp := make([][]int, length)
	for row := range dp {
		dp[row] = make([]int, length)
	}
	// 初始化临界状态dp值

	// 动态规划
	j := 0
	for l := 0; l < length; l++ {
		for i := 0; i < length; i++ {
			j = i + l
			if j >= length {
				break
			}
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 2
				} else {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1] + 2
				} else {
					dp[i][j] = maxInt(dp[i+1][j], dp[i][j-1])
				}
			}
		}
	}
	maxdp := 0
	//result := ""
	for i := 0; i < length; i++ {
		for j = 0; j < length; j++ {
			if dp[i][j] > maxdp {
				//result = s[i : j+1]
				maxdp = dp[i][j]
			}
		}
	}

	return maxdp
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
