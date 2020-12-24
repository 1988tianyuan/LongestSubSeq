package main

import "fmt"

//给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
// 注意:
//
//
// 每个数组中的元素不会超过 100
// 数组的大小不会超过 200
//
//
// 示例 1:
//
// 输入: [1, 5, 11, 5]
//
//输出: true
//
//解释: 数组可以分割成 [1, 5, 5] 和 [11].
//
//
//
//
// 示例 2:
//
// 输入: [1, 2, 3, 5]
//
//输出: false
//
//解释: 数组不能分割成两个元素和相等的子集.
//
//
//
// Related Topics 动态规划
// 👍 628 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func canPartitionNoDp(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum+=nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	half := sum>>1
	bits := make([]byte, half+1)
	bits[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := half-nums[i]; j >= 0; j-- {
			bits[j+nums[i]] |= bits[j]
			if bits[half] == 1 {
				return true
			}
		}
	}
	return false
}
func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum+=nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum/2
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1
		if i == 0 {
			for j := 0; j < target+1; j++ {
				if nums[0] == j {
					dp[0][j] = 1
				}
			}
		}
	}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		for j := 1; j < target+1; j++ {
			if num > j {
				dp[i][j] = dp[i-1][j]
			} else if num == j {
				dp[i][j] = 1
			} else {
				if dp[i-1][j] == 1 || dp[i-1][j-num] == 1 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}
			}
		}
	}
	return dp[len(nums)-1][target]==1
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartitionNoDp(nums))
	nums2 := []int{1, 2, 3, 5}
	fmt.Println(canPartitionNoDp(nums2))
	nums3 := []int{1,2,5}
	fmt.Println(canPartitionNoDp(nums3))
}
