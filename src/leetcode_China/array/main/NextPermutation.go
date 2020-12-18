package main

import "fmt"

//实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
// 必须 原地 修改，只允许使用额外常数空间。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：nums = [3,2,1]
//输出：[1,2,3]
//
//
// 示例 3：
//
//
//输入：nums = [1,1,5]
//输出：[1,5,1]
//
//
// 示例 4：
// 26586
//
//输入：nums = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100
//
// Related Topics 数组
// 👍 872 👎 0

// [2,6,5,8,6]
// [8,6,6,5,2]
// [1,3,2] -> [2,1,3]
//leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	length := len(nums)
	biggest := nums[length - 1]

	start := 0
	for i := length - 1; i >= 0; i-- {
		if nums[i] > biggest {
			biggest = nums[i]
		} else if nums[i] < biggest {
			start = i + 1
			for j := start + 1; j < length; j++ {
				if nums[j] <= nums[i] {
					swap(i, j - 1, nums)
					biggest = 0
					break
				}
			}
			if biggest != 0 {
				swap(i, length - 1, nums)
			}
			break
		}
	}
	end := length - 1
	for start < end {
		swap(start, end, nums)
		start++
		end--
	}
}

func swap(i1 int, i2 int, nums []int)  {
	tmp := nums[i2]
	nums[i2] = nums[i1]
	nums[i1] = tmp
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{2,6,5,8,6}
	nextPermutation(nums)
	fmt.Println(nums)

	nums2 := []int{8,6,6,5,2}
	nextPermutation(nums2)
	fmt.Println(nums2)

	nums3 := []int{1,3,2}
	nextPermutation(nums3)
	fmt.Println(nums3)

	nums4 := []int{1,5,1}
	nextPermutation(nums4)
	fmt.Println(nums4)

	nums5 := []int{5,4,7,5,3,2}
	nextPermutation(nums5)
	fmt.Println(nums5)
}
