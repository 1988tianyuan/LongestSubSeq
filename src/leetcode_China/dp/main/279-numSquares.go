package main

import (
	"fmt"
	"math"
)

//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//
// 示例 1:
//
// 输入: n = 12
//输出: 3
//解释: 12 = 4 + 4 + 4.
//
// 示例 2:
//
// 输入: n = 13
//输出: 2
//解释: 13 = 4 + 9.
// Related Topics 广度优先搜索 数学 动态规划
// 👍 724 👎 0

//dp[i] = min(dp(n - {1 ~ limit})) + 1
//leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	limit := int(math.Sqrt(float64(n)))
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		minVal := -1
		for j := 1; j <= limit; j++ {
			if j*j > i {
				break
			}
			if j*j == i {
				dp[i] = 1
				minVal = -1
				break
			}
			if minVal == -1 || dp[i-j*j] < minVal {
				minVal = dp[i-j*j]
			}
		}
		if minVal != -1 {
			dp[i] = minVal + 1
		}
	}
	return dp[n]
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(numSquares(14))
}
