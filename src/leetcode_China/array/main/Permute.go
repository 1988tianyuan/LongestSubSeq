package main

import (
	"fmt"
)

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
// 示例:
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
// Related Topics 回溯算法
// 👍 1043 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	return makePermute(nums, []int{}, nums, [][]int{})
}

func makePermute(nums []int, tmp []int, remain []int, result [][]int) [][]int {
	if len(remain) <= 0 {
		return append(result, tmp)
	}
	for i := 0; i < len(remain); i++ {
		current := remain[i]
		newRemain := sliceArray(i, remain)
		result = makePermute(nums, append(tmp, current), newRemain, result)
	}
	return result
}

func sliceArray(position int, remain[]int) []int {
	var newRemain []int
	for i := 0; i < len(remain); i++ {
		if i == position {
			continue
		}
		newRemain = append(newRemain, remain[i])
	}
	return newRemain
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{1,2,3}
	fmt.Println(permute(nums))
}
