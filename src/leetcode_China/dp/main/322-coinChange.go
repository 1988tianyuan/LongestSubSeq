package main

import (
	"fmt"
)

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回
// -1。
//
// 你可以认为每种硬币的数量是无限的。
//
//
//
// 示例 1：
//
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
//输入：coins = [2], amount = 3
//输出：-1
//
// 示例 3：
//
//
//输入：coins = [1], amount = 0
//输出：0
//
//
// 示例 4：
//
//
//输入：coins = [1], amount = 1
//输出：1
//
//
// 示例 5：
//
//
//输入：coins = [1], amount = 2
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 104
//
// Related Topics 动态规划
// 👍 981 👎 0

// [5,2,1]
// dp[11]=1+dp[6]=1+1+dp[1]=1+1+1=3
// = 1+dp[10]=1+1+dp[5]=3
// dp[11]=1+dp[9]=
//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount + 1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		tmp := -1
		for j := 0; j < len(coins); j++ {
			if i < coins[j] {
				continue
			}
			if dp[i-coins[j]] != -1 {
				if tmp == -1 || tmp > dp[i-coins[j]] + 1 {
					tmp = dp[i-coins[j]] + 1
				}
			}
		}
		dp[i] = tmp
	}
	return dp[amount]
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	coins := []int{1, 2, 5}
	fmt.Println(coinChange(coins, 11))
	coins2 := []int{2}
	fmt.Println(coinChange(coins2, 3))
	coins3 := []int{34,67,5,49}
	fmt.Println(coinChange(coins3, 190))
}