package main

import "fmt"

//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
//
// 示例 1 :
//
//
//输入:nums = [1,1,1], k = 2
//输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
//
//
// 说明 :
//
//
// 数组的长度为 [1, 20,000]。
// 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
//
// Related Topics 数组 哈希表
// 👍 745 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	length := len(nums)
	result := 0
	prefixSumMap := make(map[int]int)
	curSum := 0
	prefixSumMap[0] = 1
	for i:=0; i < length; i++ {
		curSum = curSum + nums[i]
		if prefixSumMap[curSum - k] != 0 {
			result += prefixSumMap[curSum - k]
		}
		prefixSumMap[curSum]++
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{1,1,1}
	fmt.Println(subarraySum(nums, 2))

	nums2 := []int{1,2,3}
	fmt.Println(subarraySum(nums2, 3))
}
