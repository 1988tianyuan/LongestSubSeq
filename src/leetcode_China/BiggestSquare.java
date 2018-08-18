package leetcode_China;

/**
 * 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
 */

public class BiggestSquare {
    public int maximalSquare(char[][] matrix) {
        if(matrix.length == 0 || matrix[0].length == 0){
            return 0;
        }

        int row = matrix.length;
        int col = matrix[0].length;
        int max = 0;
        int[][] dp = new int[row+1][col+1];
        for(int i = 1; i<=row; i++){
            for(int j = 1; j<=col; j++){
                if(matrix[i-1][j-1] == '0'){
                    dp[i][j] = 0;
                }else {
                    dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1;
                    max = dp[i][j] > max ? dp[i][j] : max;
                }
            }
        }
        return max;
    }
    private int min(int a, int b, int c){
        int min;
        if(a < b){
            min = a<c ? a : c;
        }else {
            min = b<c ? b : c;
        }
        return min;
    }
}
