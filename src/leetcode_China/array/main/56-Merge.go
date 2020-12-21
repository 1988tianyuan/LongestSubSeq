package main

import (
	"fmt"
	"sort"
)

//给出一个区间的集合，请合并所有重叠的区间。
//
//
//
// 示例 1:
//
// 输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2:
//
// 输入: intervals = [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
// 注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。
//
//
//
// 提示：
//
//
// intervals[i][0] <= intervals[i][1]
//
// Related Topics 排序 数组
// 👍 730 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type ByHead [][]int
func (intervals ByHead) Len() int {
	return len(intervals)
}
func (intervals ByHead) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}
func (intervals ByHead) Less(i, j int) bool {
	return intervals[i][0] < intervals[j][0]
}

func merge(intervals [][]int) [][]int {
	var tmp []int
	var result [][]int
	// 给intervals排序
	//{3,1,0,2}
	sort.Sort(ByHead(intervals))
	for i := 0; i < len(intervals); i++ {
		if len(tmp) == 0 {
			tmp = intervals[i]
			continue
		}
		if tmp[1] >= intervals[i][0] {
			//合并
			var start int
			var end int
			if tmp[0] > intervals[i][0] {
				start = intervals[i][0]
			} else {
				start = tmp[0]
			}
			if tmp[1] > intervals[i][1] {
				end = tmp[1]
			} else {
				end = intervals[i][1]
			}
			tmp = []int{start, end}
		} else {
			result = append(result, tmp)
			tmp = intervals[i]
		}
	}
	if len(tmp) != 0 {
		result = append(result, tmp)
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
	intervals2 := [][]int{{1,4},{4,5}}
	intervals3 := [][]int{{1,4},{0,0}}
	intervals4 := [][]int{{1,4},{0,2},{3,5}}
	fmt.Println(merge(intervals))
	fmt.Println(merge(intervals2))
	fmt.Println(merge(intervals3))
	fmt.Println(merge(intervals4))
}
