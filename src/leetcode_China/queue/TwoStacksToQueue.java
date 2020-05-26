package leetcode_China.queue;

import java.util.Stack;

/**
 * @author Liu Geng, 1988tianyuan@gmail.com
 * @date 2020/5/26 14:10
 */
public class TwoStacksToQueue {
	
	public static void main(String[] args) {
		CQueue cQueue = new CQueue();
		cQueue.appendTail(1);
		cQueue.appendTail(2);
		cQueue.appendTail(3);
		System.out.println(cQueue.deleteHead());
		System.out.println(cQueue.deleteHead());
		cQueue.appendTail(4);
		cQueue.appendTail(5);
		System.out.println(cQueue.deleteHead());
		System.out.println(cQueue.deleteHead());
		cQueue.appendTail(6);
		System.out.println(cQueue.deleteHead());
		System.out.println(cQueue.deleteHead());
		System.out.println(cQueue.deleteHead());
	}
	
	public static class CQueue {
		
		private Stack<Integer> stack1 = new Stack<>();
		private Stack<Integer> stack2 = new Stack<>();
		
		public void appendTail(int value) {
			stack1.push(value);
		}
		
		public int deleteHead() {
			if (stack2.empty()) {
				while (!stack1.empty()) {
					int value = stack1.pop();
					stack2.push(value);
				}
			}
			if (!stack2.empty()) {
				return stack2.pop();
			}
			return -1;
		}
	}
}
