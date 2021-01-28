package main

import "fmt"

//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
// 示例 1：
//
// 输入："abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//
//
// 示例 2：
//
// 输入："aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//  "cabbacb"
//	a aba b a c cbc b c    8
//
// 提示：
//
//
// 输入的字符串长度不会超过 1000 。
//
// Related Topics 字符串 动态规划
// 👍 473 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i:=0; i<len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	dp[0][0] = true
	sum := 1
	for i:=1; i<len(s); i++ {
		for j:=0; j<=i; j++ {
			if i==j {
				sum++
				dp[i][j]=true
				break
			}
			if i-j==1 {
				if s[j]==s[i] {
					sum++
					dp[i][j] = true
				}
				continue
			}
			if dp[i-1][j+1] && s[j]==s[i] {
				sum++
				dp[i][j]=true
			}
		}
	}
	return sum
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	s := "cabbacb"
	fmt.Println(countSubstrings(s))
}
