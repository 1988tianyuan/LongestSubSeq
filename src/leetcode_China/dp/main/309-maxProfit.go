package main

import "fmt"

//给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//
//
// 示例:
//
// 输入: [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
// Related Topics 动态规划
// 👍 641 👎 0
//leetcode submit region begin(Prohibit modification and deletion)
const (
	out = iota
	in
	fix
)
// {1,  2, 3, 0, 2}
// {0,  1, 2, 2, 3}	out
// {-1,-1,-1, 1, 1}	in

// {1,  2, 4}
// {0,  1, 3}
// {-1,-1,-1}

// {2,  1, 4}
// {0,  0, 3} out
// {-2,-1,-1} in
// {0,  0, 1}
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dpIn := make([]int, len(prices))
	dpIn[0] = -prices[0]
	dpOut := make([]int, len(prices))
	dpOut[0] = 0
	// 记录前一天是否out，如果前一天有out的话，当天计算dpIn就不能基于前一天的dpOut
	hasOut := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		curSell := dpIn[i-1] + prices[i]
		notSell := dpOut[i-1]
		if curSell > notSell {
			dpOut[i] = curSell
			hasOut[i] = 1
		} else {
			dpOut[i] = notSell
		}
		if hasOut[i-1] == 1 {
			if i - 2 >= 0 {
				dpIn[i] = maxProfit_309_max(dpOut[i-2] - prices[i], dpIn[i-1])
			} else {
				dpIn[i] = dpIn[i-1]
			}
		} else {
			dpIn[i] = maxProfit_309_max(dpOut[i-1] - prices[i], dpIn[i-1])
		}
	}
	return dpOut[len(dpOut)-1]
}
func maxProfit_309_max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	prices1 := []int{1,  2, 3, 0, 2}
	fmt.Println(maxProfit(prices1))
	prices2 := []int{1,  2, 4}
	fmt.Println(maxProfit(prices2))
	prices3 := []int{2,1,4}
	fmt.Println(maxProfit(prices3))
}
