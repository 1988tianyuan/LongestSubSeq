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
  if head == nil {
    return nil
  }
  return mergeSort(nil, head, nil)
}

func mergeSort(pre *SortListListNode, head *SortListListNode, tail *SortListListNode) *SortListListNode {
  if head == nil || head.Next == nil || head == tail || head.Next == tail {
    return head
  }
  newHead := head
  if head.Next.Next == tail {
    // 当前分段只剩两个元素，直接排序
    if head.Val > head.Next.Val {
      newHead = head.Next
      tmp := newHead.Next
      newHead.Next = head
      head.Next = tmp
      //if pre != nil {
      //  pre.Next = newHead
      //}
    }
    return newHead
  }
  mid := findMid(head, tail)
  tmpMidNext := mid.Next
  node1 := mergeSort(pre, head, mid.Next)
  node2 := mergeSort(mid, tmpMidNext, tail)
  tmp := node1
  if node1.Val < node2.Val {
    newHead = node1
    tmp = node1
    node1 = node1.Next
  } else {
    newHead = node2
    tmp = node2
    node2 = node2.Next
  }
  for {
    if node1 == mid.Next {
      tmp.Next = node2
      tmp = tmp.Next
      break
    }
    if node2 == tail {
      tmp.Next = node1
      tmp = tmp.Next
      break
    }
    if node1.Val < node2.Val {
      tmp.Next = node1
      tmp = node1
      node1 = node1.Next
    } else {
      tmp.Next = node2
      tmp = node2
      node2 = node2.Next
    }
  }
  return newHead
}

func findMid(head *SortListListNode, tail *SortListListNode) *SortListListNode {
  fast := head
  slow := head
  stepped := false
  for {
    fast = fast.Next
    if fast.Next == tail {
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