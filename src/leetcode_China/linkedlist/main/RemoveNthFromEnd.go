package main

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
// 示例：
//
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//
//
// 说明：
//
// 给定的 n 保证是有效的。
//
// 进阶：
//
// 你能尝试使用一趟扫描实现吗？
// Related Topics 链表 双指针
// 👍 1146 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type NthListNode struct {
  Val int
  Next *NthListNode
}

func removeNthFromEnd2(head *NthListNode, n int) *NthListNode {
	if head == nil {
		return head
	}
	if head.Next == nil && n == 1 {
		return nil
	}
	var p1 *NthListNode = nil
	p2 := head
	p2Position := 0
	for p2.Next != nil {
		p2Position++
		p2 = p2.Next
		if p2Position >= n {
			if p1 == nil {
				p1 = head
			} else {
				p1 = p1.Next
			}
		}
	}
	if p1 == nil {
		return head.Next
	}
	next := p1.Next
	if next != nil {
		// 删除p1的下一个节点
		p1.Next = next.Next
	}
	return head
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	node5 := &NthListNode{}
	node4 := &NthListNode{Next:node5}
	node3 := &NthListNode{Next:node4}
	node2 := &NthListNode{Next:node3}
	node1 := &NthListNode{Next:node2}
	removeNthFromEnd2(node1, 2)

}
