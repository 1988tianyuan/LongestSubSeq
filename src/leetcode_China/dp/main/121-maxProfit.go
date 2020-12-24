package main

import "fmt"

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
//
// 注意：你不能在买入股票前卖出股票。
//
//
//
// 示例 1:
//
// 输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//
//
// 示例 2:
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
// Related Topics 数组 动态规划
// 👍 1358 👎 0

// {7,1,5,3,6,4}
// {0,0,4,2,5,3}
//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit_121(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([]int, len(prices))
	biggest := 0
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		dp[i] = maxProfit_121_max(prices[i] - prices[i-1] + dp[i-1], 0)
		if dp[i] > biggest {
			biggest = dp[i]
		}
	}
	return biggest
}
func maxProfit_121_max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit_121(prices))
}
