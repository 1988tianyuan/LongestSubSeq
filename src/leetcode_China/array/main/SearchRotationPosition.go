package main

import "fmt"

//升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。
//
//
// 请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//
//
// 示例 2：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//
// 示例 3：
//
//
//输入：nums = [1], target = 0
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// -10^4 <= nums[i] <= 10^4
// nums 中的每个值都 独一无二
// nums 肯定会在某个点上旋转
// -10^4 <= target <= 10^4
//
// Related Topics 数组 二分查找
// 👍 1103 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	length := len(nums)
	biggest := 0
	biggestI := 0
	for i := 0; i < length; i++ {
		if nums[i] > biggest {
			biggest = nums[i]
			biggestI = i
		}
	}
	// 找到旋转点, 两边的数组用二分查找
	find1 := binarySearch(0, biggestI, nums, target)
	find2 := binarySearch(biggestI + 1, length - 1, nums, target)
	if find1 != -1 {
		return find1
	}
	if find2 != -1 {
		return find2
	}
	return -1
}

func binarySearch(start int, end int, nums []int, target int) int {
	if start == end {
		if nums[start] == target {
			return start
		}
		return -1
	} else if start > end {
		return -1
	}
	mid := (start + end) / 2
	find1 := binarySearch(start, mid, nums, target)
	find2 := binarySearch(mid + 1, end, nums, target)
	if find1 != -1 {
		return find1
	}
	if find2 != -1 {
		return find2
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{4,5,6,7,0,1,2}
	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 5))

	nums2 := []int{1,3,5}
	fmt.Println(search(nums2, 3))
}
