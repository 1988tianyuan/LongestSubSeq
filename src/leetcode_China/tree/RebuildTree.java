package leetcode_China.tree;

/**
 * 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
 * 前序遍历 preorder = [3,9,8,15,20,7,10]
 * 中序遍历 inorder = [8,9,15,3,7,20,10]
 * 	   	   3
 *    	/    \
 *    9       20
 *  /  \   	/   \
 * 8   15  7     10
 */
public class RebuildTree {
	
	public static TreeNode buildTree(int[] preorder, int[] inorder) {
		if (preorder == null || preorder.length == 0 || inorder == null || inorder.length == 0) {
			return null;
		}
		return findRoot(0, preorder.length - 1, 0, inorder.length - 1, preorder, inorder);
	}
	
	private static TreeNode findRoot(int preStart, int preEnd, int inStart, int inEnd, int[] preorder, int[] inorder) {
		int rootVal = preorder[preStart];
		TreeNode root = new TreeNode(rootVal);
		int rootIndex = 0;
		for (int i = inStart; i <= inEnd; i++) {
			if (inorder[i] == rootVal) {
				rootIndex = i;
				break;
			}
		}
		int leftCount = rootIndex - inStart;
		TreeNode leftNode = null;
		if (leftCount > 0 ){
			leftNode = findRoot(preStart + 1, preStart + leftCount, inStart, rootIndex - 1, preorder, inorder);
		}
		int rightCount = inEnd - rootIndex;
		TreeNode rightNode = null;
		if (rightCount > 0 ) {
			rightNode = findRoot(preEnd - rightCount + 1, preEnd, rootIndex + 1, inEnd, preorder, inorder);
		}
		root.left = leftNode;
		root.right = rightNode;
		return root;
	}
	
	public static void main(String[] args) {
		int[] preorder = {3,9,8,15,20,7,10};
		int[] inorder = {8,9,15,3,7,20,10};
		TreeNode root = buildTree(preorder, inorder);
		System.out.println(root);
	}
}
