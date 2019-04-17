import leetcode_China.dp.LongestIncreasSubSeq;
import leetcode_China.tree.SameTree;
import leetcode_China.tree.SymmetricTree;
import leetcode_China.tree.TreeNode;

import java.text.SimpleDateFormat;
import java.util.*;

public class Test {
    private static Random random = new Random();
    private static SimpleDateFormat simpleDateFormat = new SimpleDateFormat("hh:mm:ss");
    public static void main(String[] args){
        SymmetricTree symmetricTree = new SymmetricTree();

        TreeNode t = new TreeNode(1);
        t.left = new TreeNode(2);
        t.right = new TreeNode(2);
        t.left.left = new TreeNode(3);
        t.left.right = new TreeNode(2);
        t.right.left = new TreeNode(1);
        t.right.right = new TreeNode(3);
        System.out.println(symmetricTree.isSymmetric(t));
    }




}

