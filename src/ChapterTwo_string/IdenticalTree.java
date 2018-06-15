package ChapterTwo_string;

import ChapterTwo_string.utils.BinaryTreeNode;

public class IdenticalTree {
    public static boolean chkIdentical(BinaryTreeNode A, BinaryTreeNode B) {
        StringBuilder builderA = new StringBuilder();
        StringBuilder builderB = new StringBuilder();
        String treeStrA = serialize(A, builderA);
        String treeStrB = serialize(B, builderB);
        return KMP.kmp(treeStrA, treeStrB);
    }

    private static String serialize(BinaryTreeNode rootNode, StringBuilder builder){
        if(rootNode == null){
            return builder.append("#!").toString();
        }
        builder.append(rootNode.val+"!");
        serialize(rootNode.left, builder);
        serialize(rootNode.right, builder);
        return builder.toString();
    }
    public static void main(String[] args){
        BinaryTreeNode A = new BinaryTreeNode(1);
        BinaryTreeNode B = new BinaryTreeNode(1);
        BinaryTreeNode C = new BinaryTreeNode(1);
        BinaryTreeNode D = null;
        BinaryTreeNode E = new BinaryTreeNode(1);
        BinaryTreeNode F = new BinaryTreeNode(1);
        BinaryTreeNode G = null;
        A.left = B;
        A.right = C;
        B.left = D;
        B.right = E;
        C.left = F;
        C.right = G;

        BinaryTreeNode a = new BinaryTreeNode(1);
        BinaryTreeNode b = new BinaryTreeNode(1);
        a.right = b;

        StringBuilder builderA = new StringBuilder();
        StringBuilder builderB = new StringBuilder();
        System.out.println(serialize(A, builderA));
        System.out.println(serialize(a, builderB));

        System.out.println(chkIdentical(A, a));
    }
}


