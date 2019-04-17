package leetcode_China.tree;

/**
 * 给定一个二叉树，检查它是否是镜像对称的。
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 *     1
 *    / \
 *   2   2
 *  / \ / \
 * 3  4 4  3
 *
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 *     1
 *    / \
 *   2   2
 *    \   \
 *    3    3
 */
public class SymmetricTree {

	public boolean isSymmetric(TreeNode root) {
		if (root == null || (root.left == null && root.right == null)) {
			return true;
		}
		TreeNode left = root.left;
		TreeNode right = root.right;
		return isReverse(left, right);
	}

	private boolean isReverse(TreeNode left, TreeNode right) {
		if (left == null && right == null) {
			return true;
		} else if (left != null && right == null) {
			return false;
		} else if (left == null) {
			return false;
		}

		boolean outerReverse;
		boolean innerReverse;
		if (left.left != null && right.right != null) {
			outerReverse = isReverse(left.left, right.right);
		} else {
			outerReverse = isSame(left.left, right.right);
		}
		if (left.right != null && right.left != null) {
			innerReverse = isReverse(left.right, right.left);
		} else {
			innerReverse = isSame(left.right, right.left);
		}
		boolean rootSame = isSame(left, right);
		return rootSame && outerReverse && innerReverse;
	}

	private boolean isSame(TreeNode a, TreeNode b) {
		if (a != null && b != null) {
			return a.val == b.val;
		} else {
			return a == b;
		}
	}
}
