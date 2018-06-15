package ChapterOne_sort;

/**
 * 给定一个int矩阵mat，同时给定矩阵大小nxm及待查找的数x，请返回一个bool值，代表矩阵中是否存在x
 */
public class DoubleDimensionArrayChase {
    public static boolean findX(int[][] mat, int n, int m, int x) {
        if(mat == null || mat.length != n || mat[0].length != m){
            return false;
        }
        int indexX = n-1;
        int indexY = 0;
        while (indexX>=0 && indexY<m){
            if(mat[indexX][indexY]>x){
                indexX--;
            }else if(mat[indexX][indexY]<x){
                indexY++;
            }else {
                return true;
            }
        }
        return false;
    }
    public static void main(String[]args){
        int[][] array = new int[5][4];
        for(int i = 0; i<array.length; i++){
            for (int j = 0; j<array[i].length; j++){
                array[i][j] = i*j;
            }
        }
        for(int y = 0; y<array[0].length; y++){
            for(int x = 0; x<array.length; x++){
                System.out.print(array[x][y]+" ");
            }
            System.out.println();
        }
        System.out.println(findX(array, 5, 4, 13));
    }
}
