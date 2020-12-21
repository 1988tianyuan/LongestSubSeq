package main

import (
	"fmt"
	"sort"
)

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
// 说明：解集不能包含重复的子集。
//
// 示例:
//
// 输入: nums = [1,2,3]
//输出:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
// Related Topics 位运算 数组 回溯算法
// 👍 917 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	return sub(make([]int, 0), 0, nums, result)
}
func sub(tmp []int, start int, nums []int, result [][]int) [][]int {
	result = append(result, copySubArray(tmp))
	for i := start; i < len(nums); i++ {
		newTmp := append(tmp, nums[i])
		result = sub(newTmp, i + 1, nums, result)
	}
	return result
}
func copySubArray(array []int) []int {
	var newArray []int
	for i := 0; i < len(array); i++ {
		newArray = append(newArray, array[i])
	}
	return newArray
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{1,2,3}
	nums2 := []int{0,3,5,7,9}
	fmt.Println(subsets(nums))
	fmt.Println(subsets(nums2))
}
