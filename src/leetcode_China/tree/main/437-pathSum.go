package main

import "fmt"

//给定一个二叉树，它的每个结点都存放着一个整数值。
//
// 找出路径和等于给定数值的路径总数。
//
// 路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
// 二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
//
// 示例：
//
// root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//
//      10
//     /  \
//    5   -3
//   / \    \
//  3   2   11
// / \   \
//3  -2   1
//
//返回 3。和等于 8 的路径有:
//
//1.  5 -> 3
//2.  5 -> 2 -> 1
//3.  -3 -> 11
//
// Related Topics 树
// 👍 708 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type PathSumTreeNode struct {
  Val int
  Left *PathSumTreeNode
  Right *PathSumTreeNode
}
func pathSum(root *PathSumTreeNode, sum int) int {
	if root == nil {
		return 0
	}
	num := 0
	pathSumSearch(root, sum, &num)
	return num
}
func pathSumSearch(root *PathSumTreeNode, sum int, num *int) {
	if root == nil {
		return
	}
	pathSumInternal(root, sum, num, 0)
	pathSumSearch(root.Left, sum, num)
	pathSumSearch(root.Right, sum, num)
}
func pathSumInternal(root *PathSumTreeNode, sum int, num *int, result int) {
	curSum := result + root.Val
	if curSum == sum {
		*num++
	}
	if root.Left != nil {
		pathSumInternal(root.Left, sum, num, curSum)
	}
	if root.Right != nil {
		pathSumInternal(root.Right, sum, num, curSum)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
//      10
//     /  \
//    5   -3
//   / \    \
//  3   2   11
// / \   \
//3  -2   1
func main() {
	node9 := &PathSumTreeNode{Val:11}
	node8 := &PathSumTreeNode{Val:-3, Right:node9}
	node7 := &PathSumTreeNode{Val:1}
	node6 := &PathSumTreeNode{Val:-2}
	node5 := &PathSumTreeNode{Val:3}
	node4 := &PathSumTreeNode{Val:2, Right:node7}
	node3 := &PathSumTreeNode{Val:3, Left:node5, Right:node6}
	node2 := &PathSumTreeNode{Val:5, Left:node3, Right:node4}
	node1 := &PathSumTreeNode{Val:10, Left:node2, Right:node8}
	fmt.Println(pathSum(node1, 8))

	node21 := &PathSumTreeNode{Val:-3}
	node11 := &PathSumTreeNode{Val:-2, Right:node21}
	fmt.Println(pathSum(node11, -5))
}
