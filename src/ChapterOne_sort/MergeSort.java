package ChapterOne_sort;

import java.util.Comparator;

/**
 * 自己实现的对归并排序的简单练习，包括int数组排序和Object数组排序
 * */
public class MergeSort {

    /**
     * 以下是对整形数组进行排序
     * */
    public static void easySort(int[] array){
        //初始化设置aux的容量
        int[] aux = new int[array.length];
        easySort(array, aux, 0, array.length-1);
    }

    public static void easySort(int[]array, int[]aux, int lo, int hi){
        //如果lo=hi，则默认这个数组已经是有序的了，为递归出口
        if(hi - lo < 1){
            return;
        }

        //设置中点
        int mid = (lo + hi)>>>1;

        //对左右两部分元素分别进行排序，保证两部分数组均为有序数组
        easySort(array, aux, lo, mid);
        easySort(array, aux,mid+1, hi);

        //左右两部分均为有序数组后，进行归并
        merge(array, aux, lo, mid, hi);
    }

    //对这部分数组片段进行归并，将数组切割为左右两部分，分别进行归并（前提是左右两部分已经是有序数组了）
    public static void merge(int[]array, int[]aux, int lo, int mid, int hi){
        //将这部分数组片段复制到aux中对应的位置
        for(int k = lo; k<=hi; k++){
            aux[k] = array[k];
        }

        //设置左右两部分遍历指针起点
        int i = lo;
        int j = mid + 1;

        //对aux中的左右部分元素进行遍历并归并回array，该过程分为四种情况
        for(int k = lo; k <= hi; k++){
            if(i > mid) array[k] = aux[j++];//如果左半部分归并完了，将右半部分剩余元素扔回array
            else if(j > hi) array[k] = aux[i++];//如果右半部分归并完了，将左半部分剩余元素扔回array
            else if(aux[j]<aux[i]) array[k] = aux[j++];
            //如果左边指针元素大于右边指针元素，则将右边元素扔到array中对应位置（这个过程可以保证两元素相等时不交换原始位置），同时右边指针移动一位
            else array[k] = aux[i++];
            //同上，左边指针元素大于等于右边指针元素，则将左边元素扔到array中对应位置，同时左边指针移动一位
        }
    }

    /**
     * 以下是对引用数据类型的数组进行排序
     * */
    public static <T> void refTypeSort(T[] array, Comparator<T> c){
        //初始化设置aux的容量
        T[] aux = array.clone();
        refTypeSort(array, aux, 0, array.length-1, c);
    }

    public static <T> void refTypeSort(T[]array, T[]aux, int lo, int hi, Comparator<T> c){
        //如果lo=hi，则默认这个数组已经是有序的了，为递归出口
        if(hi - lo < 1){
            return;
        }

        //设置中点
        int mid = (lo + hi)>>>1;

        //对左右两部分元素分别进行排序，保证两部分数组均为有序数组
        refTypeSort(array, aux, lo, mid, c);
        refTypeSort(array, aux,mid+1, hi, c);

        //左右两部分均为有序数组后，进行归并
        merge(array, aux, lo, mid, hi, c);
    }


    //对这部分数组片段进行归并，将数组切割为左右两部分，分别进行归并（前提是左右两部分已经是有序数组了）
    public static <T> void merge(T[]array, T[]aux, int lo, int mid, int hi, Comparator<T> c){
        //将这部分数组片段复制到aux中对应的位置
        for(int k = lo; k<=hi; k++){
            aux[k] = array[k];
        }

        //设置左右两部分遍历指针起点
        int i = lo;
        int j = mid + 1;

        //对aux中的左右部分元素进行遍历并归并回array，该过程分为四种情况
        for(int k = lo; k <= hi; k++){
            if(i > mid) array[k] = aux[j++];//如果左半部分归并完了，将右半部分剩余元素扔回array
            else if(j > hi) array[k] = aux[i++];//如果右半部分归并完了，将左半部分剩余元素扔回array
            else if(c.compare(aux[j], aux[i])<0) array[k] = aux[j++];
                //如果左边指针元素大于右边指针元素，则将右边元素扔到array中对应位置（这个过程可以保证两元素相等时不交换原始位置），同时右边指针移动一位
            else array[k] = aux[i++];
            //同上，左边指针元素大于等于右边指针元素，则将左边元素扔到array中对应位置，同时左边指针移动一位
        }
    }
}
