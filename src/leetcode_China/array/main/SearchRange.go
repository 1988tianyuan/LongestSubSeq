package main

import (
	"fmt"
)

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 进阶：
//
//
// 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
//
//
//
//
// 示例 1：
//
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//
// 示例 2：
//
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//
// 示例 3：
//
//
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
// nums 是一个非递减数组
// -109 <= target <= 109
//
// Related Topics 数组 二分查找
// 👍 775 👎 0

// {5,7,7,8,8,8,10} -> 3,5
//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	position := searchPosition(0, len(nums) - 1, nums, target)
	if len(position) == 0 {
		return []int{-1, -1}
	}
	return position
}

func searchPosition(start int, end int, nums []int, target int) []int {
	//time.Sleep(time.Duration(100) * time.Millisecond)
	if start == end {
		if nums[start] == target {
			return []int{start, start}
		} else {
			return []int{}
		}
	} else if start > end {
		return []int{}
	}
	mid := (start + end) / 2
	var position []int
	position1 := searchPosition(start, mid, nums, target)
	position2 := searchPosition(mid + 1, end, nums, target)
	if len(position1) > 0 && len(position2) > 0 {
		return append(position, position1[0], position2[1])
	} else if len(position1) > 0 && len(position2) == 0 {
		return append(position, position1[0], position1[1])
	} else if len(position2) > 0 && len(position1) == 0 {
		return append(position, position2[0], position2[1])
	}
	return []int{}
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{5,7,7,8,8,8,10}
	fmt.Println(searchRange(nums, 8))
}
