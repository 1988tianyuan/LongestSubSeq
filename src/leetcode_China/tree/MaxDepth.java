package leetcode_China.tree;

/**
 * 给定一个二叉树，找出其最大深度
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
 * 给定二叉树 [3,9,20,null,null,15,7]
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 返回它的最大深度 3 。
 */
public class MaxDepth {
	public int maxDepth(TreeNode root) {
		int maxDepth = 0;
		if (root == null) {
			return 0;
		} else if (root.left == null && root.right != null) {
			maxDepth = maxDepth(root.right);
		} else if (root.left != null && root.right == null) {
			maxDepth = maxDepth(root.left);
		} else if (root.left != null) {
			maxDepth = Math.max(maxDepth(root.left), maxDepth(root.right));
		}
		return maxDepth + 1;
	}
}
