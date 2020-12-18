package main

import "fmt"

//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表
// 👍 1402 👎 0

//1->2->3->4->5->NULL
//5->4->3->2->1->NULL
//leetcode submit region begin(Prohibit modification and deletion)
type ReverseListNode struct {
  Val int
  Next *ReverseListNode
}
func reverseList(head *ReverseListNode) *ReverseListNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head
	var tmp1 *ReverseListNode = nil
	var tmp2 *ReverseListNode = nil
	for current != nil {
		tmp1 = current.Next
		current.Next = tmp2
		tmp2 = current
		current = tmp1
	}
	return tmp2
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	node4 := &ReverseListNode{Val:4, }
	node3 := &ReverseListNode{Val:3, Next:node4}
	node2 := &ReverseListNode{Val:2, Next:node3}
	node1 := &ReverseListNode{Val:1, Next:node2}
	reverseList(node1)
	fmt.Println(node1.Val)
}
