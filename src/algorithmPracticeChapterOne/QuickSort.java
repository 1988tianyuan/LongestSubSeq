package algorithmPracticeChapterOne;

import java.text.SimpleDateFormat;
import java.util.Date;

/**
 * 普通快速排序，选取数组最右端的元素为切分指标
 * 缺点是在初始数组随机性无法保证的情况下导致切分不平衡（即每次都将最小元素作为切分点或将最大的元素作为切分点）
 */
public class QuickSort {
    public static void quickSort(int[]array){
        quickSort(array, 0 ,array.length-1);
    }

    private static void quickSort(int[] array, int begin, int end){
        if(begin>=end){
            return;
        }
        int mid = partition(array, begin, end);
        quickSort(array, begin, mid-1);
        quickSort(array, mid+1, end);
    }
    private static int partition(int[] array, int begin, int end){
        int partiPos = end;
        int cursor = begin;
        for(int i = begin; i<end; i++){
            if(array[i]<array[partiPos]){
                swap(array, i, cursor);
                cursor++;
            }
        }
        swap(array, cursor, partiPos);
        partiPos = cursor;
        return partiPos;
    }
    private static void swap(int[]array, int i, int j){
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }

    /**
     * 三向切分快速排序，在数据重复情况较多的情况下有良好的性能
     */
    public static class QuickSortThreeWay {
        public static void quickSort(int[] array){
            quickSort(array, 0, array.length-1);
        }
        private static void quickSort(int[] array, int begin, int end){
            if(begin>=end){
                return;
            }
            int[] midDouble = partition(array, begin, end);
            quickSort(array, begin, midDouble[0]-1);
            quickSort(array, midDouble[1]+1, end);
        }
        //排序完成后，指针i和j中间的元素一定等于flag
        private static int[] partition(int[] array, int begin, int end){
            int[] midDouble = new int[2];
            int flag = array[begin];
            int i = begin;
            int j = end;
            int cursor = begin+1;
            while (cursor <= j ){
                if(array[cursor]<flag){
                    swap(array, i++, cursor++);
                }else if(array[cursor]>flag){
                    swap(array, j--, cursor);
                }else {
                    cursor++;
                }
            }
            midDouble[0] = i;
            midDouble[1] = j;
            return midDouble;
        }
    }

}









