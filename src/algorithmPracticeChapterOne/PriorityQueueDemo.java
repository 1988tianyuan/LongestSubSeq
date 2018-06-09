package algorithmPracticeChapterOne;

import java.util.AbstractQueue;
import java.util.Arrays;
import java.util.Iterator;
import java.util.Queue;

/**
 * 基于大顶堆来构建优先队列，插入和取出数据都是O(logN)的时间复杂度
 * 二叉堆：1.完全二叉树 2.每个结点都比他的两个子结点大 3.根结点一定是最大元素
 */
public class PriorityQueueDemo extends AbstractQueue<Integer>{
    private int[] array;
    private int size;

    /**
     * 构建指定初始容量的优先队列
     * @param iniCapacity
     */
    public PriorityQueueDemo(int iniCapacity){
        if(iniCapacity>=0) {
            this.array = new int[iniCapacity];
        }else {
            throw new IllegalArgumentException();
        }
    }

    /**
     * 默认初始容量16
     */
    public PriorityQueueDemo(){
        this(new int[16]);
    }

    /**
     * 指定初始数组，需要对该数组构建大顶堆
     * @param array
     */
    public PriorityQueueDemo(int[] array){
        this.array = array;
        buildMaxHeap();
    }

    public int[] getArray(){
        return array;
    }

    @Override
    public int size() {
        return size;
    }

    /**
     * 加入一个整型元素，首先放到队尾，随后通过“上浮”方式进行堆调整
     * @param integer
     * @return
     */
    @Override
    public boolean offer(Integer integer) {
        if(size+1>array.length){
            ensureCapacity(size+1);
        }
        array[size] = integer;
        swim(size);
        size++;
        return true;
    }

    /**
     * 取出队首元素，将队尾元素挪到队首，然后通过“下沉”方式进行堆调整
     * @return
     */
    @Override
    public Integer poll() {
        if(size==0){
            return null;
        }
        int result = array[0];
        swap(0 , size-1);
        array[size-1] = 0;
        sink(0);
        size--;
        return result;
    }

    @Override
    public Integer peek() {
        return null;
    }

    /**
     * 下沉调堆，当某非叶子结点改变时，从这个结点开始向下调整大顶堆
     * 一般用于取出根结点元素
     * @param i
     */
    private void sink(int i){
        if(i >= array.length){
            return;
        }
        int max = max(i, 2*i+1, 2*i+2);
        if(max != i){
            swap(i, max);
            sink(max);
        }
    }

    /**
     * 上浮调堆，当某结点改变时，从该节点开始向上调整大顶堆
     * 一般用于在队列尾部放入新元素
     * @param i
     */
    private void swim(int i){
        if(i == 0)return;
        int parent = (i-1)>>>1;
        if(array[parent]<array[i]){
            swap(i, parent);
            i = parent;
            swim(i);
        }
    }

    private void buildMaxHeap(){
        int n = array.length;
        for(int i = n/2-1; i>=0; i--){
            sink(i);
        }
    }

    private void ensureCapacity(int minCapacity){
        int oldCapacity = array.length;
        int newCapacity = oldCapacity + (oldCapacity<64 ? (oldCapacity+2) : oldCapacity>>1);
        if(newCapacity<minCapacity){
            newCapacity = minCapacity;
        }
        array = Arrays.copyOf(array, newCapacity);
    }

    private void swap(int i, int j){
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }

    private int max(int i, int j, int k){
        int n = array.length;
        if(j>=n){
            return i;
        }else if(k>=n){
            if(array[i]<array[j]){
                return j;
            }else {
                return i;
            }
        }else {
            if(array[i]<array[j]){
                if(array[j]<array[k]){
                    return k;
                }else {
                    return j;
                }
            }else if(array[i]>=array[j] && array[i]<array[k]){
                return k;
            }else {
                return i;
            }
        }
    }



    @Override
    public Iterator<Integer> iterator() {
        return null;
    }


}
