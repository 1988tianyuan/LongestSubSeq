package main

import "fmt"

//给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于
// node.val 的值之和。
//
// 提醒一下，二叉搜索树满足下列约束条件：
//
//
// 节点的左子树仅包含键 小于 节点键的节点。
// 节点的右子树仅包含键 大于 节点键的节点。
// 左右子树也必须是二叉搜索树。
//
//
// 注意：本题和 1038: https://leetcode-cn.com/problems/binary-search-tree-to-greater-s
//um-tree/ 相同
//
//
//
// 示例 1：
//
//
//
// 输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
//输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
//
//
// 示例 2：
//
// 输入：root = [0,null,1]
//输出：[1,null,1]
//
//
// 示例 3：
//
// 输入：root = [1,0,2]
//输出：[3,3,2]
//
//
// 示例 4：
//
// 输入：root = [3,2,4,1]
//输出：[7,9,4,10]
//
//
//
//
// 提示：
//
//
// 树中的节点数介于 0 和 104 之间。
// 每个节点的值介于 -104 和 104 之间。
// 树中的所有值 互不相同 。
// 给定的树为二叉搜索树。
//
// Related Topics 树
// 👍 457 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type ConvertBSTNode struct {
  Val int
  Left *ConvertBSTNode
  Right *ConvertBSTNode
}
func convertBST(root *ConvertBSTNode) *ConvertBSTNode {
	if root == nil {
		return nil
	}
	_, newRoot := convertBSTInternal(0, root)
	return newRoot
}

func convertBSTInternal(rightBigger int, cur *ConvertBSTNode) (int,*ConvertBSTNode) {
	newCur := &ConvertBSTNode{}
	rightSum := 0
	var newRight *ConvertBSTNode
	if cur.Right != nil {
		rightSum, newRight = convertBSTInternal(rightBigger, cur.Right)
	}
	newCur.Val = rightBigger + rightSum + cur.Val
	newCur.Right = newRight
	leftSum := 0
	var newLeft *ConvertBSTNode
	if cur.Left != nil {
		leftSum, newLeft = convertBSTInternal(newCur.Val, cur.Left)
	}
	newCur.Left = newLeft
	return leftSum+cur.Val+rightSum, newCur
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	node9 := &ConvertBSTNode{Val:8}
	node8 := &ConvertBSTNode{Val:7, Right:node9}
	node7 := &ConvertBSTNode{Val:3}
	node6 := &ConvertBSTNode{Val:5}
	node5 := &ConvertBSTNode{Val:6, Left:node6, Right:node8}
	node4 := &ConvertBSTNode{Val:2, Right:node7}
	node3 := &ConvertBSTNode{Val:0}
	node2 := &ConvertBSTNode{Val:1, Left:node3, Right:node4}
	node1 := &ConvertBSTNode{Val:4, Left:node2, Right:node5}
	newRoot := convertBST(node1)
	fmt.Println(newRoot.Val)
}
