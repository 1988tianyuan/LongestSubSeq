package algorithmPracticeChapterOne;

/**
 * 未采用递归的堆排序，用于判断数组是否有重复值
 * 空间复杂度O(1)
 */
public class DuplicateCheck {
    public boolean checkDuplicate(int[] a, int n) {
        if(a==null || n != a.length){
            return false;
        }
        heapSort(a);
        for(int i = 0; i<n-1; i++){
            if(a[i] == a[i+1]){
                return true;
            }
        }
        return false;
    }

    private void heapSort(int[] array){
        buildMaxHeap(array);
        int size = array.length;
        for(int i = array.length-1; i>=0; i--){
            swap(array, i, 0);
            sink(array, 0, --size);
        }
    }

    private void buildMaxHeap(int[] array){
        int n = array.length;
        for(int i = (n>>>1)-1; i>=0; i--){
            sink(array, i, n);
        }
    }

    private static void sink(int[] array, int i, int size){
        int left = 2*i+1;
        int right = 2*i+2;
        int smallest = i;
        while (i<size){
            if(left<size && array[left]<array[smallest]){
                smallest = left;
            }
            if(right<size && array[right]<array[smallest]){
                smallest = right;
            }
            if(smallest != i){
                swap(array, i, smallest);
                i = smallest;
                left = 2*i+1;
                right = 2*i+2;
            }else {
                break;
            }
        }
    }

    private static void swap(int[]array, int i, int j){
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }
}
