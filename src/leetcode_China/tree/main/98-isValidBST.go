package main

import "fmt"

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
// 示例 1:
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
//
//
// 示例 2:
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
// Related Topics 树 深度优先搜索 递归
// 👍 870 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type ValidBSTTreeNode struct {
  Val int
  Left *ValidBSTTreeNode
  Right *ValidBSTTreeNode
}
func isValidBST(root *ValidBSTTreeNode) bool {
	if root == nil {
		return false
	}
	return checkValidBST(root, -1, -1)
}
func checkValidBST(node *ValidBSTTreeNode, low int, high int) bool {
	if node == nil {
		return false
	}
	if low != -1 && node.Val <= low {
		return false
	}
	if high != -1 && node.Val >= high {
		return false
	}
	leftValid := true
	if node.Left != nil {
		if node.Left.Val >= node.Val {
			return false
		}
		leftValid = checkValidBST(node.Left, low, node.Val)
	}
	rightValid := true
	if node.Right != nil {
		if node.Right.Val <= node.Val {
			return false
		}
		rightValid = checkValidBST(node.Right, node.Val, high)
	}
	return leftValid && rightValid
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	node5 := &ValidBSTTreeNode{Val:5}
	node4 := &ValidBSTTreeNode{Val:3}
	node3 := &ValidBSTTreeNode{Val:1}
	node2 := &ValidBSTTreeNode{Val:2, Left:node3, Right:node4}
	node1 := &ValidBSTTreeNode{Val:4, Left:node2, Right:node5}
	fmt.Println(isValidBST(node1))

	node2 = &ValidBSTTreeNode{Val:1}
	node1 = &ValidBSTTreeNode{Val:1, Left:node2}
	fmt.Println(isValidBST(node1))
}
