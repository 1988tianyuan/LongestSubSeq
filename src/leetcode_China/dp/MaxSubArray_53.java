package leetcode_China.dp;

/**
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 示例:
 输入: [-2,1,-3,4,-1,2,1,-5,4],
 输出: 6
 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 */
public class MaxSubArray_53 {
    public static int maxSubArray(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        int n = nums.length;
        int[] dp = new int[n];
        int max = nums[0];
        dp[0] = nums[0];
        for (int i = 1; i < n; i++) {
            dp[i] = selectDp(dp[i-1], nums[i]);
            max = maxTwo(max, dp[i]);
        }
        return max;
    }
    private static int selectDp(int preDp, int curNum) {
        if (preDp < 0) {
            return curNum;
        }
        return preDp + curNum;
    }
    private static int maxTwo(int first, int second) {
        if (first > second) {
            return first;
        }
        return second;
    }
}
