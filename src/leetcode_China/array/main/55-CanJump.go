package main

import (
	"fmt"
)

//给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个位置。
//
// 示例 1:
//
// 输入: [2,3,1,1,4]
//输出: true
//解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
//
//
// 示例 2:
//
// 输入: [3,2,1,0,4]
//输出: false
//解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
//
// Related Topics 贪心算法 数组
// 👍 972 👎 0

//[2,3,1,1,4]
//{3,2,1,1,0}
//leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	distance := 0
	for i := 0; i<len(nums); i++ {
		cur := nums[i]
		if i > distance {
			// 超出之前节点可达范围
			return false
		}
		jumpDistance := i + cur
		if jumpDistance > distance {
			distance = jumpDistance
		}
		if distance >= len(nums) - 1 {
			// 已经可达最远
			return true
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{2,3,1,1,4}
	nums2 := []int{3,2,1,0,4}
	nums3 := []int{3,2,1,1,0}
	nums4 := []int{2,5,0,0}
	nums5 := []int{5,4,3,2,1,0}
	fmt.Println(canJump(nums))
	fmt.Println(canJump(nums2))
	fmt.Println(canJump(nums3))
	fmt.Println(canJump(nums4))
	fmt.Println(canJump(nums5))
}
