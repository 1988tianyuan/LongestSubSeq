package main

import "fmt"

//给你一个链表和一个特定值 x ，请你对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。
//
// 你应当保留两个分区中每个节点的初始相对位置。
//
//
//
// 示例：
//
//
//输入：head = 1->4->3->2->5->2, x = 3
//输出：1->2->2->4->3->5
//
// Related Topics 链表 双指针
// 👍 298 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type PartitionListNode struct {
  Val int
  Next *PartitionListNode
}
func partition(head *PartitionListNode, x int) *PartitionListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := head
	var beginNode *PartitionListNode
	var beginPre *PartitionListNode
	var curPre *PartitionListNode
	var newTmpHead *PartitionListNode
	for node != nil {
		if node.Val >= x {
			if beginNode == nil {
				beginNode = node
				beginPre = curPre
			}
		} else if beginNode != nil {
			newHead := swapPartitionNode(beginNode, beginPre, node, curPre)
			beginPre = node
			if newTmpHead == nil {
				newTmpHead = newHead
			}
		}
		curPre = node
		node = node.Next
	}
	if beginNode != nil && beginNode == head && newTmpHead != nil {
		return newTmpHead
	}
	return head
}
func swapPartitionNode(beginNode *PartitionListNode, beginPre *PartitionListNode,
	node *PartitionListNode, nodePre *PartitionListNode) *PartitionListNode {
	nodePre.Next = node.Next
	node.Next = beginNode
	if beginPre != nil {
		beginPre.Next = node
		return beginPre
	}
	return node
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	//node6 := &PartitionListNode{Val:2}
	//node5 := &PartitionListNode{Next:node6, Val:5}
	//node4 := &PartitionListNode{Next:node5, Val:2}
	//node3 := &PartitionListNode{Next:node4, Val:3}
	//node2 := &PartitionListNode{Next:node3, Val:4}
	//node1 := &PartitionListNode{Next:node2, Val:1}
	//newHead := partition(node1, 3)
	//fmt.Println(newHead.Val)

	//node91 := &PartitionListNode{Val:3}
	//node81 := &PartitionListNode{Next:node91, Val:0}
	//node71 := &PartitionListNode{Next:node81, Val:4}
	//node61 := &PartitionListNode{Next:node71, Val:1}
	//node51 := &PartitionListNode{Next:node61, Val:3}
	//node41 := &PartitionListNode{Next:node51, Val:1}
	//node31 := &PartitionListNode{Next:node41, Val:4}
	//node21 := &PartitionListNode{Next:node31, Val:0}
	//node11 := &PartitionListNode{Next:node21, Val:2}
	//newHead1 := partition(node11, 4)
	//fmt.Println(newHead1.Val)

	node22 := &PartitionListNode{Val:1}
	node12 := &PartitionListNode{Next:node22, Val:1}
	newHead2 := partition(node12, 0)
	fmt.Println(newHead2.Val)
}
