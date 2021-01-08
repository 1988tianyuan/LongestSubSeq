package main

import (
	"container/list"
	"fmt"
)

//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]
//
//
//
//
//
// 示例 1:
//
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//输出: 3
//解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
//
//
// 示例 2:
//
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
//输出: 5
//解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
//
//
//
//
// 说明:
//
//
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉树中。
//
// Related Topics 树
// 👍 891 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type LCATreeNode struct {
	Val int
	Left *LCATreeNode
	Right *LCATreeNode
}
func lowestCommonAncestor(root, p, q *LCATreeNode) *LCATreeNode {
	if root == p || root == q {
		return root
	}
	treeStackP := list.New()
	treeStackQ := list.New()
	if !findNode(root.Left, p, treeStackP) {
		findNode(root.Right, p, treeStackP)
	}
	if !findNode(root.Left, q, treeStackQ) {
		findNode(root.Right, q, treeStackQ)
	}
	lastSame := root
	for treeStackP.Len() != 0 && treeStackQ.Len() != 0 {
		eP := treeStackP.Back()
		eQ := treeStackQ.Back()
		if eP.Value != eQ.Value {
			break
		} else {
			lastSame = eP.Value.(*LCATreeNode)
			treeStackP.Remove(eP)
			treeStackQ.Remove(eQ)
		}
	}
	return lastSame
}
func findNode(node *LCATreeNode, target *LCATreeNode, treeStack *list.List) bool {
	if node == nil {
		return false
	}
	if node == target {
		treeStack.PushBack(node)
		return true
	} else {
		leftR := findNode(node.Left, target, treeStack)
		rightR := findNode(node.Right, target, treeStack)
		if leftR || rightR {
			treeStack.PushBack(node)
		}
		return leftR || rightR
	}
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	node8 := &LCATreeNode{Val:8}
	node7 := &LCATreeNode{Val:7}
	node6 := &LCATreeNode{Val:6}
	node5 := &LCATreeNode{Val:5, Left:node6, Right:node7}
	node4 := &LCATreeNode{Val:4, Right:node8}
	node3 := &LCATreeNode{Val:3}
	node2 := &LCATreeNode{Val:2, Left:node3, Right:node4}
	node1 := &LCATreeNode{Val:1, Left:node2, Right:node5}
	fmt.Println(lowestCommonAncestor(node1, node2, node8).Val)
}
