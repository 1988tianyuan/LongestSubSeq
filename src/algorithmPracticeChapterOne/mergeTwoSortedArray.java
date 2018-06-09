package algorithmPracticeChapterOne;

/**
 * 有两个从小到大排序以后的数组A和B，其中A的末端有足够的缓冲空容纳B。请编写一个方法，将B合并入A并排序。
 * 给定两个有序int数组A和B，A中的缓冲空用0填充，同时给定A和B的真实大小int n和int m，返回合并后的数组。
 */
public class mergeTwoSortedArray {
    public static int[] mergeAB(int[] A, int[] B, int n, int m){
        int indexA = n-1;
        int indexB = m-1;
        for(int i = n+m-1; i>=0; i--){
            if(indexA<0){
                A[i] = B[indexB];
                indexB--;
            }else if(indexB<0){
                A[i] = A[indexA];
                indexA--;
            }else if(A[indexA]>B[indexB]){
                A[i] = A[indexA];
                indexA--;
            }else {
                A[i] = B[indexB];
                indexB--;
            }
        }
        return A;
    }
}
