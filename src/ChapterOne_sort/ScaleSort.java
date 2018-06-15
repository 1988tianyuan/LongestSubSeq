package ChapterOne_sort;

/**
 * 改良版堆排序，适用于数组元素位置与有序后位置差值不大于k的情况
 * 时间复杂度O(nlogk)，空间复杂度O(k)
 */
public class ScaleSort {
    public int[] sortElement(int[] A, int n, int k) {
        if(n != A.length || k>n || A==null){
            return A;
        }
        //初始化大小为k的暂存小顶堆数组，其元素为A数组中前k个元素
        int[] bufferHeap = buildInitMinHeap(A, k);

        //遍历前n-k个元素，将暂存数组顶部元素弹回到原始数组，保证原始数组0到n-k-1的元素是有序的
        for(int i = k; i<n; i++){
            A[i-k] = bufferHeap[0]; //将暂存堆顶部元素拷贝到原始数组i-k的位置
            bufferHeap[0] = A[i];   //将后续元素复制到暂存堆顶部
            heapify(bufferHeap, 0, k);  //顶部元素改变，由上往下调整堆，使之保持小顶堆结构
        }

        //还剩下n-k至n-1的元素是无序的，对这部分元素采用常规堆排序方式进行排序
        for(int i = n - k; i<n; i++){
            A[i] = bufferHeap[0];   //堆顶元素拷贝回原始数组
            swap(bufferHeap, 0 , k-1);  //交换暂存堆的首尾元素
            heapify(bufferHeap, 0, --k);    //堆大小减一，由上至下调堆
        }
        return A;
    }

    //构建大小为k的暂存小顶堆数组，并初始化构建这个小顶堆
    private int[] buildInitMinHeap(int[] array, int k){
        int[] bufferHeap = new int[k];
        for(int i = 0; i<k; i++){
            insertHeap(bufferHeap, i, array[i]);
        }
        return bufferHeap;
    }

    //构建暂存小顶堆数组时，将原始数组前K个值一一插入到暂存数组尾部
    //插入到尾部后，需要通过swim的方式至底向上调整堆结构
    private void insertHeap(int[] bufferHeap, int i, int value){
        bufferHeap[i] = value;
        while (i>0){
            int parent = (i-1)>>>1;
            if(bufferHeap[parent]>bufferHeap[i]){
                swap(bufferHeap, i, parent);
                i = parent;
            }else {
                break;
            }
        }
    }

    //调整堆，当堆顶元素改变时通过sink方式至顶向下调整堆结构
    private void heapify(int[] array, int i, int heapSize){
        int left = 2*i+1;
        int right = 2*i+2;
        int smallest = i;
        if(left<heapSize && array[left]<array[smallest]){
            smallest = left;
        }
        if(right<heapSize && array[right]<array[smallest]){
            smallest = right;
        }
        if(smallest != i){
            swap(array, i, smallest);
            heapify(array, smallest, heapSize);
        }
    }

    private void swap(int[]array, int i, int j){
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }
}
