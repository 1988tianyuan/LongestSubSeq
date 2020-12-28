package main

import "fmt"

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例：
//二叉树：[3,9,20,null,null,15,7],
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层序遍历结果：
//
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]
//
// Related Topics 树 广度优先搜索
// 👍 733 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type LevelOrderTreeNode struct {
  Val int
  Left *LevelOrderTreeNode
  Right *LevelOrderTreeNode
}
func levelOrder(root *LevelOrderTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	level := 0
	return levelTraversal(root, result, level)
}
func levelTraversal(node *LevelOrderTreeNode, result [][]int, level int) [][]int {
	if node == nil {
		return result
	}
	if level >= len(result) {
		result = append(result, []int{})
	}
	result[level] = append(result[level], node.Val)
	result = levelTraversal(node.Left, result, level+1)
	result = levelTraversal(node.Right, result, level+1)
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	node5 := &LevelOrderTreeNode{Val:5}
	node4 := &LevelOrderTreeNode{Val:3}
	node3 := &LevelOrderTreeNode{Val:1}
	node2 := &LevelOrderTreeNode{Val:2, Left:node3, Right:node4}
	node1 := &LevelOrderTreeNode{Val:4, Left:node2, Right:node5}
	fmt.Println(levelOrder(node1))
}
