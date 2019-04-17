package leetcode_China.tree;

/**
 * 给定两个二叉树，编写一个函数来检验它们是否相同。
 * 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 *
 * 输入:       1         1
 *           / \       / \
 *          2   3     2   3
 *
 *         [1,2,3],   [1,2,3]
 *
 * 输出: true
 */
public class SameTree {

	public boolean isSameTree(TreeNode p, TreeNode q) {
		if (p == null && q == null) {
			return true;
		} else if (p != null && q == null) {
			return false;
		} else if (p == null) {
			return false;
		}

		boolean leftSame;
		boolean rightSame;
		boolean rootSame;
		if (p.left != null && q.left != null) {
			leftSame = isSameTree(p.left, q.left);
		} else {
			leftSame = isSame(p.left, q.left);
		}
		if (p.right != null && q.right != null) {
			rightSame = isSameTree(p.right, q.right);
		} else {
			rightSame = isSame(p.right, q.right);
		}
		rootSame = isSame(p, q);
		return rootSame && leftSame && rightSame;
	}

	private boolean isSame(TreeNode a, TreeNode b) {
		if (a != null && b != null) {
			return a.val == b.val;
		} else {
			return a == b;
		}
	}

}
