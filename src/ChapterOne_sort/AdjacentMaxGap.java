package ChapterOne_sort;

/**
 * 有一个整形数组A，请设计一个复杂度为O(n)的算法，算出排序后相邻两数的最大差值。
 * 给定一个int数组A和A的大小n，请返回最大的差值
 */
public class AdjacentMaxGap {
    public static int maxGap(int[] A, int n) {
        if(A == null || A.length != n){
            return 0;
        }
        int max = A[0];
        int min = A[0];
        for(int i = 0; i<n; i++){
            if(A[i]>max){
                max = A[i];
            }else if(A[i]<min){
                min = A[i];
            }
        }
        if(min == max)return 0;
        //初始化一个大小为n+1的数组用于标记每个桶是否有元素，由于最大元素单独存放于最后一个桶，因此一共n+1个桶
        boolean[] bucketHasNums = new boolean[n+1];
        //初始化两个数组用于存放每个桶的最大值和最小值
        int[] maxs = new int[n+1];
        int[] mins = new int[n+1];
        //遍历所有元素，找到当前元素归属的桶下标，并配置当前桶的最大和最小值，最后将该桶设置为“有元素”
        for(int i = 0; i<n; i++){
            int bid = bucket(A[i], n , max, min);   //计算当前元素应该归属的桶编号
            maxs[bid] = bucketHasNums[bid] ? Math.max(maxs[bid], A[i]) : A[i];
            mins[bid] = bucketHasNums[bid] ? Math.min(mins[bid], A[i]) : A[i];
            bucketHasNums[bid] = true;
        }
        int maxGap = 0; //保存相邻元素最大差值
        int index = 0;  //遍历每个桶的指针
        int lastMax = 0;   //当前桶的上一个非空桶的最大值
        while (index<n+1){
            //找到第一个非空桶
            if(bucketHasNums[index]){
                lastMax = maxs[index];
                break;
            }
            index++;
        }
        //从第二个非空桶开始遍历，计算当前桶最小值和上一个非空桶最大值的差值，并更新结果
        for(int i = index; i<n+1; i++){
            if(bucketHasNums[i]){
                maxGap = Math.max(maxGap, mins[i]-lastMax);
                lastMax = maxs[i];
            }
        }
        return maxGap;
    }
    //判断某数num是属于哪一个桶的，返回桶数组的下标
    private static int bucket(long num, long n, long max, long min){
        return (int)((num - min)*n/(max - min));
    }
}
