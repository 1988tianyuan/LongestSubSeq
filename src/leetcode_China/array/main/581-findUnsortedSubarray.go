package main

import "fmt"

//给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
// 你找到的子数组应是最短的，请输出它的长度。
//
// 示例 1:
//
//
//输入: [2, 6, 4, 8, 10, 9, 15]
//输出: 5
//解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//
//
// 说明 :
//
//
// 输入的数组长度范围在 [1, 10,000]。
// 输入的数组可能包含重复元素 ，所以升序的意思是<=。
//
// Related Topics 数组
// 👍 459 👎 0

// {2, 6, 4, 8, 10, 9, 15}
//leetcode submit region begin(Prohibit modification and deletion)
func findUnsortedSubarrayTwoLoop(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	max := nums[0]
	begin := 0
	end := 0
	subRangeLen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}
	}
	min := nums[end]
	for j := end - 1; j >=0; j-- {
		if nums[j] < min {
			min = nums[j]
		}  else if nums[j] > min {
			begin = j
		}
	}
	if end != begin {
		subRangeLen = end - begin + 1
	} else {
		subRangeLen = 0
	}
	return subRangeLen
}
func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	max := nums[0]
	begin := len(nums)
	subRangeLen := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num >= max {
			max = num
		} else {
			var tmpBegin int
			if begin == len(nums) {
				tmpBegin = i - 2
			} else {
				tmpBegin = begin - 1
			}
			for tmpBegin >= 0 && nums[tmpBegin] > num {
				tmpBegin--
			}
			begin = findUnsortedSubarrayMin(tmpBegin + 1, begin)
			subRangeLen = i - begin + 1
		}
	}
	return subRangeLen
}
func findUnsortedSubarrayMin(i1 int, i2 int) int {
	if i1 > i2 {
		return i2
	}
	return i1
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarrayTwoLoop(nums))
	nums2 := []int{2, 1}
	fmt.Println(findUnsortedSubarrayTwoLoop(nums2))
	nums3 := []int{2,3,3,2,4}
	fmt.Println(findUnsortedSubarrayTwoLoop(nums3))
	nums4 := []int{1,2,4,5,3}
	fmt.Println(findUnsortedSubarrayTwoLoop(nums4))
	nums5 := []int{1,3,5,4,2}
	fmt.Println(findUnsortedSubarrayTwoLoop(nums5))
	nums6 := []int{2,1,1,1,1}
	fmt.Println(findUnsortedSubarrayTwoLoop(nums6))
	nums7 := []int{1,3,2,4,5}
	fmt.Println(findUnsortedSubarrayTwoLoop(nums7))
}
