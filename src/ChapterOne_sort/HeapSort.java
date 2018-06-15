package ChapterOne_sort;

/**
 * 堆排序，首先构造大顶堆，将根结点元素和尾部元素交换，再将剩余元素进行调堆
 * 将交换后的尾部构成有序序列，依次循环，直到剩余元素全部有序
 * 加强版的选择排序，时间复杂度O(nlogn)，其中建堆时间复杂度为O(n)
 * 在建堆过程中，相同元素前后顺序可能被打乱，所以是不稳定的排序
 * 由于采用了递归，占用了额外的栈空间，空间复杂度O为(logn)
 */
public class HeapSort {
    public static void heapSort(int[] array){
        buildMaxHeap(array);
        int size = array.length;
        for(int i = array.length-1; i>=0; i--){
            swap(array, i, 0);
            sink(array, 0, --size);
        }
    }

    private static void buildMaxHeap(int[] array){
        int n = array.length;
        for(int i = (n>>>1)-1; i>=0; i--){
            sink(array, i, n);
        }
    }

    private static void sink(int[] array, int i, int size){
        int max = max(array, i, 2*i+1, 2*i+2, size);
        if(max != i){
            swap(array, i, max);
            sink(array, max, size);
        }
    }

    private static int max(int[] array, int i, int j, int k, int size){
        int n = size;
        if(j>=n){
            return i;
        }else if(k>=n){
            if(array[i]<array[j]){
                return j;
            }
            return i;
        }else {
            if(array[i]<array[j]){
                if(array[j]<array[k]){
                    return k;
                }
                return j;
            }else if(array[i]>=array[j] && array[i]<array[k]){
                return k;
            }
            return i;
        }
    }

    private static void swap(int[]array, int i, int j){
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }
}


