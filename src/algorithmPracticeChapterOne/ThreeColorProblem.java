package algorithmPracticeChapterOne;

/**
 * 给定一个只含0，1，2的整数数组A及它的大小，请返回排序后的数组。保证数组大小小于等于500
 * 只能使用原地排序
 */
public class ThreeColorProblem {
    public static int[] sortThreeColor(int[] A, int n) {
        if(A==null || A.length != n){
            return A;
        }
        int indexLeft = 0;
        int indexRight = n-1;
        int i = 0;
        while (i<=indexRight){
            if(A[i] == 0){
                swap(A, indexLeft++, i);
            }else if(A[i] == 2){
                swap(A, indexRight--, i);
                continue;
            }
            i++;
        }
        return A;
    }

    private static void swap(int[]array, int i, int j){
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }
}
