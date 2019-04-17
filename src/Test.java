import leetcode_China.dp.LongestIncreasSubSeq;
import leetcode_China.tree.InvertTree;
import leetcode_China.tree.MaxDepth;
import leetcode_China.tree.SameTree;
import leetcode_China.tree.SymmetricTree;
import leetcode_China.tree.TreeNode;

import java.text.SimpleDateFormat;
import java.util.*;

public class Test {
    private static Random random = new Random();
    private static SimpleDateFormat simpleDateFormat = new SimpleDateFormat("hh:mm:ss");
    public static void main(String[] args){
        MaxDepth maxDepth = new MaxDepth();
        TreeNode t = new TreeNode(3);
        t.left = new TreeNode(9);
        System.out.println(maxDepth.maxDepth(t));
    }




}

