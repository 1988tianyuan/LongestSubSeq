package main

import "container/list"

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

	}
}
//leetcode submit region end(Prohibit modification and deletion)
