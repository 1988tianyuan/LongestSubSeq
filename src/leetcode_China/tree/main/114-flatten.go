package main

import "fmt"

//给定一个二叉树，原地将它展开为一个单链表。
//
//
//
// 例如，给定二叉树
//
//     1
//   / \
//  2   5
// / \   \
//3   4   6
//
// 将其展开为：
//
// 1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
// Related Topics 树 深度优先搜索
// 👍 671 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type FlattenTreeNode struct {
  Val int
  Left *FlattenTreeNode
  Right *FlattenTreeNode
}
func flatten(root *FlattenTreeNode)  {
	flattenNode(root)
}
func flattenNode(node *FlattenTreeNode) *FlattenTreeNode {
	if node == nil {
		return nil
	}
	var end *FlattenTreeNode
	var leftEnd *FlattenTreeNode
	var rightEnd *FlattenTreeNode
	if node.Left != nil {
		leftEnd = flattenNode(node.Left)
	}
	if node.Right != nil {
		rightEnd = flattenNode(node.Right)
	}
	if leftEnd != nil {
		leftEnd.Right = node.Right
		end = leftEnd
	}
	if rightEnd != nil {
		end = rightEnd
	}
	if end == nil {
		end = node
	}
	if node.Left != nil {
		node.Right = node.Left
		node.Left = nil
	}
	return end
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	node6 := &FlattenTreeNode{Val:6}
	node5 := &FlattenTreeNode{Val:5, Right:node6}
	node4 := &FlattenTreeNode{Val:4}
	node3 := &FlattenTreeNode{Val:3}
	node2 := &FlattenTreeNode{Val:2, Left:node3, Right:node4}
	node1 := &FlattenTreeNode{Val:1, Left:node2, Right:node5}
	flatten(node1)
	fmt.Println(node1.Val)

	node3 = &FlattenTreeNode{Val:3}
	node2 = &FlattenTreeNode{Val:2, Right:node3}
	node1 = &FlattenTreeNode{Val:1, Left:node2}
	flatten(node1)
	fmt.Println(node1.Val)

	node3 = &FlattenTreeNode{Val:3}
	node2 = &FlattenTreeNode{Val:2, Left:node3}
	node1 = &FlattenTreeNode{Val:1, Left:node2}
	flatten(node1)
	fmt.Println(node1.Val)

	node3 = &FlattenTreeNode{Val:3}
	node2 = &FlattenTreeNode{Val:2}
	node1 = &FlattenTreeNode{Val:1, Left:node2, Right:node3}
	flatten(node1)
	fmt.Println(node1.Val)
}
