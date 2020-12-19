package main

import (
	"fmt"
	"sort"
)

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的数字可以无限制重复被选取。
//
// 说明：
//
//
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
//
//
// 示例 1：
//
// 输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//  [7],
//  [2,2,3]
//]
//
//
// 示例 2：
//
// 输入：candidates = [2,3,5], target = 8,
//所求解集为：
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//]
//
//
//
// 提示：
//
//
// 1 <= candidates.length <= 30
// 1 <= candidates[i] <= 200
// candidate 中的每个元素都是独一无二的。
// 1 <= target <= 500
//
// Related Topics 数组 回溯算法
// 👍 1093 👎 0

// {2,3,5}
//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	var result [][]int
	sort.Ints(candidates)
	return combineSum(candidates, []int{}, 0, result, target)
}

func combineSum(candidates []int, tmp []int, start int, result [][]int, target int) [][]int {
	if start >= len(candidates) {
		return result
	}
	sum := 0
	for i := 0; i < len(tmp); i++ {
		sum += tmp[i]
	}
	for i := start; i < len(candidates); i++ {
		if sum + candidates[i] > target {
			break
		} else if sum + candidates[i] == target {
			// 必须拷贝一个新数组
			newnewTmp := append(copyArray(tmp), candidates[i])
			result = append(result, newnewTmp)
		} else {
			newnewTmp := append(copyArray(tmp), candidates[i])
			result = combineSum(candidates, newnewTmp, i, result, target)
		}
	}
	return result
}

func copyArray(array []int) []int {
	var newArray []int
	for i := 0; i < len(array); i++ {
		newArray = append(newArray, array[i])
	}
	return newArray
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	candidates := []int{2,3,5}
	candidates2 := []int{2,3,6,7}
	candidates3 := []int{1}
	candidates4 := []int{2,7,6,3,5,1}
	fmt.Println(combinationSum(candidates, 8))
	fmt.Println(combinationSum(candidates2, 7))
	fmt.Println(combinationSum(candidates3, 2))
	fmt.Println(combinationSum(candidates4, 9))
}
