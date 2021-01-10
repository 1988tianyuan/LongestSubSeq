package main

import "fmt"

//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。
//这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“
//房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
//如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
// 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
//
// 示例 1:
//
// 输入: [3,2,3,null,3,null,1]
//
//     3
//    / \
//   2   3
//    \   \
//     3   1
//
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
//
// 示例 2:
//
// 输入: [3,4,5,1,3,null,1]
//
//     3
//    / \
//   4   5
//  / \   \
// 1   3   1
//
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
//
// Related Topics 树 深度优先搜索
// 👍 692 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type HouseRobberTreeNode struct {
  Val int
  Left *HouseRobberTreeNode
  Right *HouseRobberTreeNode
}
func rob(root *HouseRobberTreeNode) int {
  if root == nil {
    return 0
  }
  val1, val2 := robInternal(root)
  if val1 > val2 {
    return val1
  }
  return val2
}
func robInternal(root *HouseRobberTreeNode) (int, int) {
  if root.Left == nil && root.Right == nil {
    return root.Val, 0
  }
  var leftUseRootVal int
  var leftNotUseRootVal int
  if root.Left != nil {
    leftUseRootVal, leftNotUseRootVal = robInternal(root.Left)
  }
  var rightUseRootVal int
  var rightNotUseRootVal int
  if root.Right != nil {
    rightUseRootVal, rightNotUseRootVal = robInternal(root.Right)
  }
  useRootVal := robberMax(root.Val + rightNotUseRootVal + leftNotUseRootVal,
    leftUseRootVal + rightUseRootVal)
  return useRootVal, leftUseRootVal + rightUseRootVal
}
func robberMax(a1,a2 int) int {
  if a1 > a2 {
    return a1
  }
  return a2
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
  node8 := &HouseRobberTreeNode{Val:8}
  node7 := &HouseRobberTreeNode{Val:7}
  node6 := &HouseRobberTreeNode{Val:6}
  node5 := &HouseRobberTreeNode{Val:5, Left:node6, Right:node7}
  node4 := &HouseRobberTreeNode{Val:4, Right:node8}
  node3 := &HouseRobberTreeNode{Val:3}
  node2 := &HouseRobberTreeNode{Val:2, Left:node3, Right:node4}
  node1 := &HouseRobberTreeNode{Val:1, Left:node2, Right:node5}
  fmt.Println(rob(node1))

  node11 := &HouseRobberTreeNode{Val:1}
  node12 := &HouseRobberTreeNode{Val:3}
  node13 := &HouseRobberTreeNode{Val:3, Right:node11}
  node14 := &HouseRobberTreeNode{Val:2, Right:node12}
  node15 := &HouseRobberTreeNode{Val:3, Left:node14, Right:node13}
  fmt.Println(rob(node15))

  node21 := &HouseRobberTreeNode{Val:3}
  node22 := &HouseRobberTreeNode{Val:2, Right:node21}
  node23 := &HouseRobberTreeNode{Val:1, Right:node22}
  node24 := &HouseRobberTreeNode{Val:4, Left:node23}
  fmt.Println(rob(node24))
}
