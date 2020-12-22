package main

import (
	"container/list"
	"fmt"
)

//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,2,3]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树
// 👍 481 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type PreorderTraversalTreeNode struct {
  Val int
  Left *PreorderTraversalTreeNode
  Right *PreorderTraversalTreeNode
}
func preorderTraversal(root *PreorderTraversalTreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := list.New()
	node := root
	for {
		for node != nil {
			result = append(result, node.Val)
			stack.PushFront(node)
			node = node.Left
		}
		if stack.Len() == 0 {
			break
		}
		nodeE := stack.Front()
		node = nodeE.Value.(*PreorderTraversalTreeNode)
		stack.Remove(nodeE)
		node = node.Right
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	node5 := &PreorderTraversalTreeNode{Val:5}
	node4 := &PreorderTraversalTreeNode{Val:4}
	node3 := &PreorderTraversalTreeNode{Val:3}
	node2 := &PreorderTraversalTreeNode{Val:2, Left:node3, Right:node4}
	node1 := &PreorderTraversalTreeNode{Val:1, Left:node2, Right:node5}
	fmt.Println(preorderTraversal(node1))
}