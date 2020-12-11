package main

import "fmt"

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//
//
// 示例 2：
//
// 输入: "cbbd"
//输出: "bb"
//
// Related Topics 字符串 动态规划
// 👍 2981 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	length := len(s)
	var dp = make([][]bool, length)
	for i:= range dp {
		dp[i] = make([]bool, length)
	}
	start := 0
	end := 0
	longest := 0
	for i := 1; i < length; i++ {
		dp[i][i] = true
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] && (i - j < 3 || dp[i-1][j+1]) {
				dp[i][j] = true
				if i - j + 1 > longest {
					start = j
					end = i
					longest = end - start + 1
				}
			}
		}
	}
	return s[start:end+1]
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	s := "aacabdkacaa"
	fmt.Printf(longestPalindrome(s))
}
