package main

import "fmt"

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
// Related Topics 深度优先搜索 广度优先搜索 并查集
// 👍 927 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	islandNodeMap := make([][]*IslandNode, len(grid))
	for i := 0; i < len(islandNodeMap); i++ {
		islandNodeMap[i] = make([]*IslandNode, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			islandNodeMap[i][j] = &IslandNode{}
		}
	}
	rootMap := make(map[*IslandNode]byte)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			curNode := islandNodeMap[i][j]
			var upNode *IslandNode
			if i!=0 && grid[i-1][j] == 1 {
				upNode = islandNodeMap[i-1][j]
			}
			var leftNode *IslandNode
			if j!= 0 && grid[i][j-1] == 1 {
				leftNode = islandNodeMap[i][j-1]
			}
			newRoot := unionNode(upNode, leftNode)
			if newRoot != nil {
				curNode.father = newRoot
			}
		}
	}
	nums := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			root := islandNodeMap[i][j]
			for root.father != nil {
				root = root.father
			}
			if rootMap[root] != 1 {
				nums++
				rootMap[root] = 1
			}
		}
	}
	return nums
}

func unionNode(node1 *IslandNode, node2 *IslandNode) *IslandNode {
	var node1Root *IslandNode
	node1Deep := 0
	if node1 != nil {
		node1Root = node1.father
		if node1Root == nil {
			node1Root = node1
		} else {
			node1Deep = 1
			for node1Root.father != nil {
				node1Root = node1Root.father
				node1Deep++
			}
		}
	}
	var node2Root *IslandNode
	node2Deep := 0
	if node2 != nil {
		node2Root = node2.father
		if node2Root == nil {
			node2Root = node2
		} else {
			node2Deep = 1
			for node2Root.father != nil {
				node2Root = node2Root.father
				node2Deep++
			}
		}
	}
	if node1Root != nil && node2Root != nil {
		if node1Root == node2Root {
			return node1Root
		}
		if node1Deep > node2Deep {
			node2Root.father = node1Root
			return node1Root
		} else {
			node1Root.father = node2Root
			return node2Root
		}
	}
	if node1Root != nil && node2Root == nil {
		return node1Root
	}
	if node1Root == nil && node2Root != nil {
		return node2Root
	}
	return nil
}

type IslandNode struct {
	father *IslandNode
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	//grid := [][]byte{
	// {1,1,0,0,0},
	// {1,1,0,0,0},
	// {0,0,1,0,0},
	// {0,0,0,1,1},
	//}
	//fmt.Println(numIslands(grid))
	grid2 := [][]byte{
		{1,1,0,1,1},
		{1,1,1,1,1},
		{0,0,0,0,0},
		{1,1,1,1,1},
		{0,1,0,1,1},
	}
	fmt.Println(numIslands(grid2))
	grid3 := [][]byte{
		{1,1,0,1,1},
		{1,1,1,1,1},
		{0,1,1,1,0},
		{1,1,1,1,1},
		{0,1,0,1,1},
	}
	fmt.Println(numIslands(grid3))
	grid4 := [][]byte{
		{1,1,1,1,0},
		{1,1,0,1,0},
		{1,1,0,0,0},
		{0,0,0,0,0},
	}
	fmt.Println(numIslands(grid4))
}
