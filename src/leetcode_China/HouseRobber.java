package leetcode_China;

/**
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
 * 影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
 * 输入: [1,2,3,1]
 * 输出: 4
 * 解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 */
public class HouseRobber {

    public int rob(int[] nums) {
        if (nums == null || nums.length == 0) return 0;
        int dp0 = 0;
        int dp1 = 0;
        int temp;
        for (int num : nums) {
            temp = dp1;
            dp1 = Math.max(dp0 + num, dp1);
            dp0 = temp;
        }
        return Math.max(dp0, dp1);
    }
}
