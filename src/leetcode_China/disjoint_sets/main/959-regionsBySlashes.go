package main

import "fmt"

//在由 1 x 1 方格组成的 N x N 网格 grid 中，每个 1 x 1 方块由 /、\ 或空格构成。这些字符会将方块划分为一些共边的区域。
//
// （请注意，反斜杠字符是转义的，因此 \ 用 "\\" 表示。）。
//
// 返回区域的数目。
//
//
//
//
//
//
// 示例 1：
//
// 输入：
//[
//  " /",
//  "/ "
//]
//输出：2
//解释：2x2 网格如下：
//
//
// 示例 2：
//
// 输入：
//[
//  " /",
//  "  "
//]
//输出：1
//解释：2x2 网格如下：
//
//
// 示例 3：
//
// 输入：
//[
//  "\\/",
//  "/\\"
//]
//输出：4
//解释：（回想一下，因为 \ 字符是转义的，所以 "\\/" 表示 \/，而 "/\\" 表示 /\。）
//2x2 网格如下：
//
//
// 示例 4：
//
// 输入：
//[
//  "/\\",
//  "\\/"
//]
//输出：5
//解释：（回想一下，因为 \ 字符是转义的，所以 "/\\" 表示 /\，而 "\\/" 表示 \/。）
//2x2 网格如下：
//
//
// 示例 5：
//
// 输入：
//[
//  "//",
//  "/ "
//]
//输出：3
//解释：2x2 网格如下：
//
//
//
//
//
// 提示：
//
//
// 1 <= grid.length == grid[0].length <= 30
// grid[i][j] 是 '/'、'\'、或 ' '。
//
// Related Topics 深度优先搜索 并查集 图
// 👍 190 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
const (
	SPACE = 32
	FORWARD_SLASH = 47
	BACK_SLASH = 92

	leftUp = 1
	rightDown = 2
	leftDown = 3
	rightUp = 4
	allSpace = 5
)

func regionsBySlashes(grid []string) int {
	sideLen := len(grid)
	//初始化并查集
	sets := make([]*SlashNode, sideLen * sideLen * 2)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			node := grid[i][j]
			sets = addNode(node, sets, sideLen*i + j)
		}
	}
	// 合并并查集
	for i := 1; i < len(sets); i++ {
		curNode := sets[i]
		if curNode.slash == rightDown {
			continue
		}
		if curNode.slash != rightUp && i%(sideLen*2) != 0 {
			// 和左边节点合并
			unionSlashNode(i-1, i, sets)
		}
		if curNode.slash != leftDown && i-2*sideLen >= 0 {
			// 和上方节点合并
			upOriginI := i/2-sideLen
			up0Index := upOriginI*2
			up1Index := upOriginI*2 + 1
			up0 := sets[up0Index]
			if up0.slash == leftDown || up0.slash == rightDown || up0.slash == allSpace {
				unionSlashNode(up0Index, i, sets)
			} else {
				unionSlashNode(up1Index, i, sets)
			}
		}
	}
	// 遍历并查集数组找到root总数
	rootMap := make(map[int]int)
	for i := 0; i < len(sets); i++ {
		rootI := findRoot(i, sets)
		rootMap[rootI] = 1
	}
	return len(rootMap)
}

func unionSlashNode(node1Index int, node2Index int, sets []*SlashNode) {
	root1Index := findRoot(node1Index, sets)
	root2Index := findRoot(node2Index, sets)
	if root1Index != root2Index {
		sets[root2Index].fatherIndex = root1Index
	}
}

func findRoot(i int, sets []*SlashNode) int {
	node := sets[i]
	fatherI := node.fatherIndex
	fatherNode := sets[fatherI]
	for fatherI != i {
		fatherI = fatherNode.fatherIndex
		fatherNode = sets[fatherI]
		i = fatherI
	}
	return fatherI
}


func addNode(slash uint8, sets []*SlashNode, i int) []*SlashNode {
	node1Slash := 0
	node2Slash := 0
	switch slash {
	case FORWARD_SLASH:
		node1Slash = leftUp
		node2Slash = rightDown
		break
	case BACK_SLASH:
		node1Slash = leftDown
		node2Slash = rightUp
		break
	case SPACE:
		node1Slash = allSpace
		node2Slash = allSpace
	}
	node1 := &SlashNode{slash:node1Slash, fatherIndex:2*i}
	sets[2*i] = node1
	node2 := &SlashNode{slash:node2Slash, fatherIndex:2*i+1}
	sets[2*i+1] = node2
	return sets
}
type SlashNode struct {
	slash int
	fatherIndex int
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	//// 3
	//grid := []string{"//","/ "}
	//fmt.Println(regionsBySlashes(grid))
	//
	//// 5
	//grid2 := []string{"/\\","\\/"}
	//fmt.Println(regionsBySlashes(grid2))
	//
	//// 4
	//grid3 := []string{"\\/","/\\"}
	//fmt.Println(regionsBySlashes(grid3))
	//
	//// 2
	//grid4 := []string{" /","/ "}
	//fmt.Println(regionsBySlashes(grid4))
	//
	//// 1
	//grid5 := []string{" /","  "}
	//fmt.Println(regionsBySlashes(grid5))
	//
	//// 4
	//grid6 := []string{" /\\"," \\/","\\  "}
	//fmt.Println(regionsBySlashes(grid6))

	// 3
	grid7 := []string{"\\/\\ "," /\\/"," \\/ ","/ / "}
	fmt.Println(regionsBySlashes(grid7))
}
