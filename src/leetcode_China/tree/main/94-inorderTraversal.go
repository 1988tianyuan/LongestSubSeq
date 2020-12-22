package main

import (
	"container/list"
	"fmt"
)

//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
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
//输出：[2,1]
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
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表
// 👍 804 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	return traversalNoRecursion(root, result)
}
func traversal(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}
	left := node.Left
	right := node.Right
	result = traversal(left, result)
	result = append(result, node.Val)
	result = traversal(right, result)
	return result
}

func traversalNoRecursion(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}
	stack := list.New()
	for {
		for node != nil {
			stack.PushFront(node)
			node = node.Left
		}
		if stack.Len() == 0 {
			break
		}
		nodeE := stack.Front()
		node = nodeE.Value.(*TreeNode)
		result = append(result, node.Val)
		stack.Remove(nodeE)
		node = node.Right
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	node5 := &TreeNode{Val:5}
	node4 := &TreeNode{Val:4}
	node3 := &TreeNode{Val:3}
	node2 := &TreeNode{Val:2, Left:node3, Right:node4}
	node1 := &TreeNode{Val:1, Left:node2, Right:node5}
	fmt.Println(inorderTraversal(node1))
}
