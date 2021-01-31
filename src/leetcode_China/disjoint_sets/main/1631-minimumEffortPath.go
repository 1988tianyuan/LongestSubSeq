package main

import (
	"fmt"
	"sort"
)

//你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子 (row,
// col) 的高度。一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从 0 开始编号）。你
//每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。
//
// 一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。
//
// 请你返回从左上角走到右下角的最小 体力消耗值 。
//
//
//
// 示例 1：
//
//
//
//
//输入：heights = [[1,2,2],[3,8,2],[5,3,5]]
//输出：2
//解释：路径 [1,3,5,3,5] 连续格子的差值绝对值最大为 2 。
//这条路径比路径 [1,2,2,2,5] 更优，因为另一条路径差值最大值为 3 。
//
//
// 示例 2：
//
//
//
//
//输入：heights = [[1,2,3],[3,8,4],[5,3,5]]
//输出：1
//解释：路径 [1,2,3,4,5] 的相邻格子差值绝对值最大为 1 ，比路径 [1,3,5,3,5] 更优。
//
//
// 示例 3：
//
//
//输入：heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
//输出：0
//解释：上图所示路径不需要消耗任何体力。
//
//
//
//
// 提示：
//
//
// rows == heights.length
// columns == heights[i].length
// 1 <= rows, columns <= 100
// 1 <= heights[i][j] <= 106
//
// Related Topics 深度优先搜索 并查集 图 二分查找
// 👍 80 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func minimumEffortPath(heights [][]int) int {
	if len(heights) == 1 && len(heights[0]) == 1 {
		return 0
	}
	lines := make([]*EffortLine, 0)
	effortSet := make([][][]int, len(heights))
	for i:=0; i<len(heights); i++ {
		effortSet[i] = make([][]int, len(heights[0]))
		for j:=0; j<len(heights[i]); j++ {
			if j != len(heights[i])-1 {
				cur := &EffortLine{beginX:j, beginY:i, endX:j+1, endY:i}
				if heights[i][j] - heights[i][j+1] >= 0 {
					cur.value = heights[i][j] - heights[i][j+1]
				} else {
					cur.value = heights[i][j+1] - heights[i][j]
				}
				lines = append(lines, cur)
			}
			if i != len(heights)-1 {
				cur := &EffortLine{beginX:j, beginY:i, endX:j, endY:i+1}
				if heights[i][j] - heights[i+1][j] >= 0 {
					cur.value = heights[i][j] - heights[i+1][j]
				} else {
					cur.value = heights[i+1][j] - heights[i][j]
				}
				lines = append(lines, cur)
			}
			effortSet[i][j] = []int{i, j}
		}
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i].value - lines[j].value <= 0
	})
	for i:=0; i<len(lines); i++ {
		cur := lines[i]
		beginRoot, beginDepth := findEffortRoot(cur.beginX, cur.beginY, effortSet, true)
		endRoot, endDepth := findEffortRoot(cur.endX, cur.endY, effortSet, true)
		if beginRoot[0] != endRoot[0] || beginRoot[1] != endRoot[1] {
			// 合并
			if beginDepth > endDepth {
				effortSet[endRoot[0]][endRoot[1]] = beginRoot
			} else {
				effortSet[beginRoot[0]][beginRoot[1]] = endRoot
			}
			// 判断首尾是否相连
			headRoot, _ := findEffortRoot(0, 0, effortSet, false)
			tailRoot, _ := findEffortRoot(len(heights[0])-1, len(heights)-1, effortSet, false)
			if headRoot[0] == tailRoot[0] && headRoot[1] == tailRoot[1] {
				return lines[i].value
			}
		}
	}
	return lines[len(lines)-1].value
}

func findEffortRoot(x int, y int, effortSet [][][]int, flat bool) ([]int,int) {
	curX := x
	curY := y
	father := effortSet[y][x]
	depth := 0
	for father[0] != curY || father[1] != curX {
		depth++
		curY = father[0]
		curX = father[1]
		father = effortSet[father[0]][father[1]]
	}
	// 扁平化
	if flat {
		cur := effortSet[y][x]
		for cur[0] != y || cur[1] != x {
			y = cur[0]
			x = cur[1]
			cur = effortSet[y][x]
			effortSet[y][x] = father
		}
	}
	return father,depth
}

type EffortLine struct {
	beginX int
	beginY int
	endX int
	endY int
	value int
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	heights := [][]int{
		{1,2,2},{3,8,2},{5,3,5},
	}
	fmt.Println(minimumEffortPath(heights))
	heights2 := [][]int{
		{1,2,3},{3,8,4},{5,3,5},
	}
	fmt.Println(minimumEffortPath(heights2))
	heights3 := [][]int{
		{1,2,1,1,1},
		{1,2,1,2,1},
		{1,2,1,2,1},
		{1,2,1,2,1},
		{1,1,1,2,1},
	}
	fmt.Println(minimumEffortPath(heights3))
	heights4 := [][]int{
		{3},
	}
	fmt.Println(minimumEffortPath(heights4))
	heights5 := [][]int{
		{1,10,6,7,9,10,4,9},
	}
	fmt.Println(minimumEffortPath(heights5))
}
