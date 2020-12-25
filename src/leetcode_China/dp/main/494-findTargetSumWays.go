package main

import (
	"fmt"
)

//给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选
//择一个符号添加在前面。
//
// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
//
//
//
// 示例：
//
// 输入：nums: [1, 1, 1, 1, 1], S: 3
//输出：5
//解释：
//
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
//
//一共有5种方法让最终目标和为3。
//
//
//
//
// 提示：
//
//
// 数组非空，且长度不会超过 20 。
// 初始的数组的和不会超过 1000 。
// 保证返回的最终结果能被 32 位整数存下。
//
// Related Topics 深度优先搜索 动态规划
// 👍 504 👎 0

// {1, 1, 1, 1, 1} 1
// {
//  {1,1},
//	{1,2},	dp[i][j] = dp[i-1][j] + dp[i][j-nums[i]]
//	{1,3},
//	{1,4},
//	{1,5},
// }
//leetcode submit region begin(Prohibit modification and deletion)
func findTargetDp(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum+=nums[i]
	}
	target := (sum-S)/2
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1
	}
	if nums[0] <= target {
		dp[0][nums[0]] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j < target+1; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-nums[i]]
		}
	}
	return dp[len(nums)-1][target]
}

func findTargetDfs(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	return findTarget(nums, 0, 0, 0, S)
}
func findTarget(nums []int, tmp int, position int, result int, target int) int {
	if position >= len(nums) {
		if tmp == target {
			result++
		}
		return result
	}
	newTmp := tmp + nums[position]
	result = findTarget(nums, newTmp, position+1, result, target)
	newTmp = tmp - nums[position]
	result = findTarget(nums, newTmp, position+1, result, target)
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetDp(nums, 3))
	nums2 := []int{1,0}
	fmt.Println(findTargetDp(nums2, 1))
}
