package main

import (
	"container/list"
	"fmt"
)

//请判断一个链表是否为回文链表。
//
// 示例 1:
//
// 输入: 1->2
//输出: false
//
// 示例 2:
//
// 输入: 1->2->2->1
//输出: true
// 1->3->2->2->3->1
//
//
// 进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
// Related Topics 链表 双指针
// 👍 796 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type PalindromeListNode struct {
  Val int
  Next *PalindromeListNode
}

//最优解
func isPalindrome2(head *PalindromeListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast := head
	slow := head
	slowStep := false

	for fast != nil {
		fast = fast.Next
		if slowStep {
			slow = slow.Next
			slowStep = false
		} else {
			slowStep = true
		}
	}
	fast = head

	// 反转链表
	var tmp2 *PalindromeListNode = nil
	var tmp1 *PalindromeListNode = nil
	for slow != nil {
		tmp1 = slow.Next
		slow.Next = tmp2
		tmp2 = slow
		slow = tmp1
	}
	slow = tmp2

	for fast != nil && slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

func isPalindrome(head *PalindromeListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	queue := list.New()
	current := head
	second := head
	length := 0
	for current != nil {
		current = current.Next
		length++
	}
	isOdd := false
	if length % 2 != 0 {
		isOdd = true
	}
	half := length / 2
	if half == 1 {
		isPal := head.Val == head.Next.Val
		if !isPal && head.Next.Next != nil {
			isPal = head.Val == head.Next.Next.Val
		}
		return isPal
	}
	for half > 0 {
		queue.PushBack(second.Val)
		second = second.Next
		half--
	}
	if isOdd {
		second = second.Next
	}
	for second != nil {
		e := queue.Back()
		queue.Remove(e)
		if e.Value.(int) != second.Val {
			return false
		}
		second = second.Next
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	node5 := &PalindromeListNode{Val:1}
	node4 := &PalindromeListNode{Next:node5,Val:3}
	//node3 := &PalindromeListNode{Next:node4,Val:2}
	node2 := &PalindromeListNode{Next:node4,Val:2}
	node1 := &PalindromeListNode{Next:node2,Val:1}
	fmt.Println(isPalindrome2(node1))
}
