package main

import "fmt"

//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
// 进阶：
//
//
// 你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
//
//
//
// 示例 1：
//
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
//
// 示例 2：
//
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 5 * 104] 内
// -105 <= Node.val <= 105
//
// Related Topics 排序 链表
// 👍 935 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type SortListListNode struct {
  Val int
  Next *SortListListNode
}

func sortList(head *SortListListNode) *SortListListNode {
  if head == nil || head.Next == nil {
    return head
  }
  var newHead *SortListListNode
  mid := findMid(head)
  tmpMidNext := mid.Next
  mid.Next = nil
  node1 := sortList(head)
  node2 := sortList(tmpMidNext)
  var tmp *SortListListNode
  for {
    if node1 == nil {
      if tmp == nil {
        newHead = node2
      } else {
        tmp.Next = node2
      }
      break
    }
    if node2 == nil {
      if tmp == nil {
        newHead = node1
      } else {
        tmp.Next = node1
      }
      break
    }
    if node1.Val < node2.Val {
      if tmp == nil {
        newHead = node1
      } else {
        tmp.Next = node1
      }
      tmp = node1
      node1 = node1.Next
    } else {
      if tmp == nil {
        newHead = node2
      } else {
        tmp.Next = node2
      }
      tmp = node2
      node2 = node2.Next
    }
  }
  return newHead
}

func findMid(head *SortListListNode) *SortListListNode {
  fast := head
  slow := head
  stepped := false
  for {
    fast = fast.Next
    if fast.Next == nil {
      break
    }
    if stepped {
      stepped = false
    } else {
      slow = slow.Next
      stepped = true
    }
  }
  return slow
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
  node4 := &SortListListNode{Val:2}
  node3 := &SortListListNode{Next:node4, Val:4}
  node2 := &SortListListNode{Next:node3, Val:1}
  node1 := &SortListListNode{Next:node2, Val:3}
  newHead := sortList(node1)
  fmt.Println(newHead.Val)
}