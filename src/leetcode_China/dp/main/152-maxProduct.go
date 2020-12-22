package main

import "fmt"

//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//
//
// 示例 1:
//
// 输入: [2,3,-1,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//
//
// 示例 2:
//
// 输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
// Related Topics 数组 动态规划
// 👍 870 👎 0

// {2,3,-1,4,-2,4}
// {2,6, 6,6, 8,32}
// {2,2,-3,-4,-8,-32}

// {2,3,-2,2,4}
// {2,6,-2,2,8}
// {2,2,-6,-24,-96}

// {2,3,0,4,2}
// {2,6,0,4,8}
// {2,2,0,0,0}

// {-2,3,-4}
// {-2,3,24}
// {-2,-6,-6}
//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dpNeg := make([]int, len(nums))
	biggest := nums[0]
	dp[0] = nums[0]
	dpNeg[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = maxProductMax(dp[i-1] * nums[i], dpNeg[i-1] * nums[i], nums[i])
		dpNeg[i] = minProductMax(dp[i-1] * nums[i], dpNeg[i-1] * nums[i], nums[i])
		if dp[i] > biggest {
			biggest = dp[i]
		}
	}
	return biggest
}
func maxProductMax(i1 int, i2 int, i3 int) int {
	if i1 > i2 {
		if i1 > i3 {
			return i1
		} else {
			return i3
		}
	} else {
		if i2 > i3 {
			return i2
		} else {
			return i3
		}
	}
}
func minProductMax(i1 int, i2 int, i3 int) int {
	if i1 < i2 {
		if i1 < i3 {
			return i1
		} else {
			return i3
		}
	} else {
		if i2 < i3 {
			return i2
		} else {
			return i3
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{2,3,-1,4,-2,4}
	fmt.Println(maxProduct(nums))
	nums2 := []int{2,3,-2,2,4}
	fmt.Println(maxProduct(nums2))
	nums3 := []int{2,3,-2,4}
	fmt.Println(maxProduct(nums3))
	nums4 := []int{-2,3,-4}
	fmt.Println(maxProduct(nums4))
}
