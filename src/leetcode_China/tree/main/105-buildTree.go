package main

import "fmt"

//根据一棵树的前序遍历与中序遍历构造二叉树。
//
// 注意:
//你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 前序遍历 preorder = [3,9,4,8,20,15,7]
//中序遍历 inorder = [4,9,8,3,15,20,7]
//
// 返回如下的二叉树：
//
//    3
//   / \
//  9  20
// /\  / \
//4  8 15 7
// Related Topics 树 深度优先搜索 数组
// 👍 807 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type BuildTreeTreeNode struct {
  Val int
  Left *BuildTreeTreeNode
  Right *BuildTreeTreeNode
}
// preorder = [3,9,4,8,20,15,7]
// inorder = [4,9,8,3,15,20,7]
//    3
//   / \
//  9  20
// /\  / \
//4  8 15 7
var numMap map[int]int
func buildTree(preorder []int, inorder []int) *BuildTreeTreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	numMap = make(map[int]int, len(preorder))
	for i := 0; i < len(inorder); i++ {
		numMap[inorder[i]] = i
	}
	return buildTreeWithIndex(preorder, inorder, 0, len(preorder)-1, 0)
}
func buildTreeWithIndex(preorder []int, inorder []int, preStart int, preEnd int, inStart int) *BuildTreeTreeNode {
	if preStart >= len(preorder) || preEnd >= len(preorder) || inStart >= len(inorder) || preStart > preEnd {
		return nil
	}
	totalLen := preEnd-preStart+1
	root := &BuildTreeTreeNode{Val:preorder[preStart]}
	rootI := numMap[root.Val]
	leftTreeLen := rootI - inStart
	root.Left = buildTreeWithIndex(preorder, inorder, preStart+1, preStart+leftTreeLen, inStart)
	root.Right = buildTreeWithIndex(preorder, inorder, preStart+leftTreeLen+1, preStart+totalLen-1, rootI+1)
	return root
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	preorder := []int{3,9,4,8,20,15,7}
	inorder := []int{4,9,8,3,15,20,7}
	node := buildTree(preorder, inorder)
	fmt.Println(node.Val)
}
