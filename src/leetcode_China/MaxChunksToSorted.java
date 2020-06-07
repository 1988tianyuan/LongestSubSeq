package leetcode_China;

/**
 * #769
 * 数组arr是[0, 1, ..., arr.length - 1]的一种排列，我们将这个数组分割成几个“块”，并将这些块分别进行排序。
 * 之后再连接起来，使得连接的结果和按升序排序后的原数组相同。我们最多能将数组分成多少块？
 * 输入: arr = [4,3,2,1,0] [1,0,4,2,3]
 * 输出: 1
 */
public class MaxChunksToSorted {

    public static int maxChunksToSorted(int[] arr) {
        int max = 0;
        int num = 0;
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] > max) {
                max = arr[i];
            }
            if (max == i) {
                num++;
            }
        }
        return num;
    }

    public static void main(String[] args) {
        int[] arr = {1,0,2,4,3};
        System.out.println(maxChunksToSorted(arr));
    }
}
