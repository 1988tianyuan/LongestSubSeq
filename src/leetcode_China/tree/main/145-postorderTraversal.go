package main

import (
	"container/list"
	"fmt"
)

//给定一个二叉树，返回它的 后序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [3,2,1]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树
// 👍 493 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type PostorderTraversalTreeNode struct {
	Val int
	Left *PostorderTraversalTreeNode
	Right *PostorderTraversalTreeNode
}
func postorderTraversal(root *PostorderTraversalTreeNode) []int {
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
			node = node.Right
		}
		if stack.Len() == 0 {
			break
		}
		nodeE := stack.Front()
		node = nodeE.Value.(*PostorderTraversalTreeNode)
		stack.Remove(nodeE)
		node = node.Left
	}
	start := 0
	end := len(result) - 1
	for start < end {
		result[start], result[end] = result[end], result[start]
		start++
		end--
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	node5 := &PostorderTraversalTreeNode{Val:5}
	node4 := &PostorderTraversalTreeNode{Val:4}
	node3 := &PostorderTraversalTreeNode{Val:3}
	node2 := &PostorderTraversalTreeNode{Val:2, Left:node3, Right:node4}
	node1 := &PostorderTraversalTreeNode{Val:1, Left:node2, Right:node5}
	fmt.Println(postorderTraversal(node1))
}