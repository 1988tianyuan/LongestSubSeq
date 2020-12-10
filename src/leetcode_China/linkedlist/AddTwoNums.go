package linkedlist

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
// Related Topics 链表 数学
// 👍 5359 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var curVal int
	isStep := false
	var realVal int
	if l1 == nil && l2 != nil {
		curVal = l2.Val
	} else if l1 != nil && l2 == nil {
		curVal = l1.Val
	} else if l1 != nil && l2 != nil {
		curVal = l1.Val + l2.Val
	} else {
		return nil
	}
	if curVal >= 10 {
		isStep = true
		realVal = curVal - 10
	} else {
		realVal = curVal
	}
	var nextNode *ListNode
	if isStep {
		if l1 != nil && l2 != nil {
			if l1.Next != nil {
				l1.Next.Val = l1.Next.Val + 1
			} else if l2.Next != nil {
				l2.Next.Val = l2.Next.Val + 1
			} else {
				l1.Next = &ListNode{Val:1}
			}
		} else if l1 != nil {
			if l1.Next != nil {
				l1.Next.Val = l1.Next.Val + 1
			} else {
				l1.Next = &ListNode{Val:1}
			}
		} else if l2 != nil {
			if l2.Next != nil {
				l2.Next.Val = l2.Next.Val + 1
			} else {
				l2.Next = &ListNode{Val:1}
			}
		}
	}

	if l1 == nil && l2 != nil {
		if l2.Next == nil {
			nextNode = nil
		} else {
			nextNode = addTwoNumbers(nil, l2.Next)
		}
	} else if l1 != nil && l2 == nil {
		if l1.Next == nil {
			nextNode = nil
		} else {
			nextNode = addTwoNumbers(l1.Next, nil)
		}
	} else if l1 != nil && l2 != nil {
		nextNode = addTwoNumbers(l1.Next, l2.Next)
	}
	var curNode *ListNode
	if nextNode == nil {
		curNode = &ListNode{Val:realVal}
	} else {
		curNode = &ListNode{Val:realVal, Next:nextNode}
	}
	return curNode
}
//leetcode submit region end(Prohibit modification and deletion)
