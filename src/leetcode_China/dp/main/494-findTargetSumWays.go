package main

import "fmt"

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
// 👍 505 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWaysDpOptimization(nums []int, S int) int {
	if len(nums) == 0 {
		if S == 0 {
			return 1
		}
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == S || nums[0] == -S {
			return 1
		}
		return 0
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum+=nums[i]
	}
	if (sum - S) < 0 || (sum - S)&1 == 1 {
		return 0
	}
	target := (sum-S)/2
	dp := make([]int, target+1)
	if nums[0] <= target {
		dp[nums[0]] = 1
	}
	if nums[0] == 0 {
		dp[0] = 2
	} else {
		dp[0] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j-nums[i]] + dp[j]
		}
	}
	return dp[target]
}

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		if S == 0 {
			return 1
		}
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == S || nums[0] == -S {
			return 1
		}
		return 0
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum+=nums[i]
	}
	if (sum - S) < 0 || (sum - S)&1 == 1 {
		return 0
	}
	target := (sum-S)/2
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, target+1)
	}
	if nums[0] <= target {
		dp[0][nums[0]] = 1
	}
	if nums[0] == 0 {
		dp[0][0] = 2
	} else {
		dp[0][0] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < target+1; j++ {
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-nums[i]] + dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}
func findTargetSumWays_min(i1 int, i2 int) int {
	if i1 > i2 {
		return i2
	}
	return i1
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{1,2,2,2,3}
	fmt.Println(findTargetSumWaysDpOptimization(nums, 8))
	nums2 := []int{0,1}
	fmt.Println(findTargetSumWaysDpOptimization(nums2, 1))
	nums3 := []int{1000}
	fmt.Println(findTargetSumWaysDpOptimization(nums3, -1000))
	nums4 := []int{0,0,0,0,0,0,0,0,1}
	fmt.Println(findTargetSumWaysDpOptimization(nums4, 1))
	nums5 := []int{1,1,1,1,1}
	fmt.Println(findTargetSumWaysDpOptimization(nums5, 3))
}