package algorithmPracticeChapterOne;

/**
 * 计算需要排序的最短子数组的长度
 */
public class SubSeqNeedToSort {
    public static int shortestSubsequence(int[] A, int n) {
        if(A == null || A.length != n){
            return 0;
        }
        int leftMax = A[0]; //记录从左边开始遍历过的元素中最大值
        int indexLeftMax = 0;  //若当前元素小于目前的最大值时进行标记
        //从左向右遍历
        for(int i = 0; i<n; i++){
            if(A[i]>leftMax){
                leftMax = A[i];
            }
            if(A[i]<leftMax){
                indexLeftMax = i;
            }
        }
        int rightMin = A[n-1];  //记录从右边开始遍历过的元素中最小值
        int indexRightMin = n-1;   //若当前元素大于目前的最小值时进行标记
        //从右向左遍历
        for(int j = n-1; j>=0; j--){
            if(A[j]<rightMin){
                rightMin = A[j];
            }
            if(A[j]>rightMin){
                indexRightMin = j;
            }
        }
        //从右往左遍历和从左往右遍历的最终标记点中间的元素就是需要进行排序的最短子数组
        return (indexLeftMax - indexRightMin)>0 ? (indexLeftMax - indexRightMin)+1 : 0;
    }
}
