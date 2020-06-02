package leetcode_China.linkedlist;

/**
 * 将两个升序链表合并为一个新的升序链表并返回, 新链表是通过拼接给定的两个链表的所有节点组成的。 
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 */
public class MergeTwoLists {

    public static ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode tmp;
        if (l1 != null && l2 != null) {
            if (l1.val < l2.val) {
                tmp = l1.next;
                l1.next = mergeTwoLists(tmp, l2);
                return l1;
            } else {
                tmp = l2.next;
                l2.next = mergeTwoLists(tmp, l1);
                return l2;
            }
        } else {
            return l1 != null ? l1 : l2;
        }
    }

    public static void main(String[] args) {
        ListNode l1 = new ListNode(1);
        l1.next = new ListNode(2);
        l1.next.next = new ListNode(4);
        ListNode l2 = new ListNode(1);
        l2.next = new ListNode(3);
        l2.next.next = new ListNode(4);
        ListNode head = mergeTwoLists(l1, l2);
        System.out.println(head);
    }
}
