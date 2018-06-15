package ChapterOne_sort;

/**
 * 希尔排序时间复杂度取决于步长gap的选择，在O(n3/2)—O(n2)之间，空间复杂度为O(1)
 * 各步长的分步排序过程将打乱之前元素的先后顺序，所以是不稳定的
 * 总体性能优于简单插入排序
 */
public class ShellSort {
    public static int[] shellSort(int[]A, int n){
        if(A == null || A.length != n){
            return A;
        }
        //初始gap为n/2，此后每次减半
        for(int gap = n>>>1; gap>0; gap=gap>>>1){
            //对每个gap进行简单插入排序
            for(int i = gap; i<n; i++){
                int j = i;
                while (j-gap>=0 && A[j]<A[j-gap]){
                    swap(A, j ,j-gap);
                    j-=gap;
                }
            }
        }
        return A;
    }

    private static void swap(int[]array, int i, int j){
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }
}
