package leetcode_China;

/**
 * 给定一个矩阵 A， 返回 A 的转置矩阵。
 矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
 */
public class MatrixTranspose {
    public int[][] transpose(int[][] A) {
        int[][] B = new int[A[0].length][A.length];
        for(int i = 0; i<A.length; i++){
            for(int j = 0; j<A[0].length; j++){
                B[j][i] = A[i][j];
            }
        }
        return B;
    }
}
