package main

import (
	"fmt"
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
// [-4, -1, -1, 0, 1, 2]
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针
// 👍 2816 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	var mid int
	var end int
	for start := 0; start < len(nums); start++ {
		if nums[start] > 0 {
			// 起始大于0直接退出
			break
		}
		if start - 1 >= 0 && nums[start - 1] == nums[start] {
			// 跳过
			continue
		}
		mid = start + 1
		end = len(nums) - 1
		for mid < end {
			value := nums[start] + nums[mid] + nums[end]
			if value < 0 {
				// mid右移
				mid++
			} else if value == 0 {
				result = append(result, []int{nums[start], nums[mid], nums[end]})
				mid++
				end--
			} else {
				// end左移
				end--
			}
			for mid != start + 1 && mid < end && mid - 1 >= 0 && nums[mid-1] == nums[mid] {
				// 跳过
				mid++
			}
			for end + 1 < len(nums) && end >= 0 && nums[end+1] == nums[end] {
				end--
			}
		}
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Printf("%v", result)
	nums = []int{0,0,0,0}
	result = threeSum(nums)
	fmt.Printf("%v", result)
	result = threeSum(nil)
	fmt.Printf("%v", result)
	result = threeSum([]int{-1,0,1,2,-1,-4,-2,-3,3,0,4})
	fmt.Printf("%v", result)
}
