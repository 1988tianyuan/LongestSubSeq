package leetcode_China.dp;

/**
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 *
 * 示例:
 *
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 * 4,10,4,3,8,9
 * 10,2,5,7,101
 * [1,3,6,7,9,4,10,5,6]
 */
public class LongestIncreasSubSeq {

	public int lengthOfLIS(int[] nums) {
		int len;
		if (nums == null || (len = nums.length) == 0) {
			return 0;
		}
		// dp数组，dp[i]是以nums[i]为结尾的最大上升子序列的长度
		// 如果nums=[4,10,4,3,8,9]，则dp=[1,2,1,1,2,3]
		int[] dp = new int[len];
		int longest = 0;
		for (int i = 0; i < len; i++) {
			dp[i] = 1;
			for (int j = 0; j < i; j++) {
				if (nums[j] < nums[i]) {
					dp[i] = Math.max(dp[i], dp[j] + 1);
				}
			}
			longest = Math.max(dp[i], longest);
		}
		return longest;
	}
}
