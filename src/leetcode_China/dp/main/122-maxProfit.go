package main

import "fmt"

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1:
//
// 输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
//
//
// 示例 2:
//
// 输入: [1,2,3,4,5]
//输出: 4
//解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//
//
// 示例 3:
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 3 * 10 ^ 4
// 0 <= prices[i] <= 10 ^ 4
//
// Related Topics 贪心算法 数组
// 👍 1039 👎 0

// {7,  1, 5, 3, 6, 4}
// {0,  0, 4, 2, 7, 5}	卖出
// {-7,-1,-1, 1, 1, 3}	买入
//leetcode submit region begin(Prohibit modification and deletion)
const (
	buyIn = iota
	sellOut
)
func maxProfit_122(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	biggest := 0
	dp[0][buyIn] = -prices[0]
	dp[0][sellOut] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][sellOut] = maxProfit_122_max(dp[i-1][buyIn] + prices[i], dp[i-1][sellOut])
		dp[i][buyIn] = maxProfit_122_max(dp[i-1][sellOut] - prices[i], dp[i-1][buyIn])
		biggest = maxProfit_122_max(biggest, dp[i][sellOut])
	}
	return biggest
}
func maxProfit_122_max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
//leetcode submit region end(Prohibit modification and deletion)
// {3,  2, 6, 5, 0, 3}
// {0,  0, 4, 4, 4, 7}	卖出
// {-3,-2,-2,-1, 4, 4}	买入
func main() {
	prices := []int{7,  1, 5, 3, 6, 4}
	fmt.Println(maxProfit_122(prices))
	prices2 := []int{1,2,3,4,5}
	fmt.Println(maxProfit_122(prices2))
	prices3 := []int{3,2,6,5,0,3}
	prices4 := []int{1,  2, 3, 0, 2}
	prices5 := []int{1,  2, 4}
	fmt.Println(maxProfit_122(prices3))
	fmt.Println(maxProfit_122(prices4))
	fmt.Println(maxProfit_122(prices5))
}
