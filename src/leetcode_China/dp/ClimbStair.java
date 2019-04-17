package leetcode_China.dp;

public class ClimbStair {

	public int climbStairs(int n) {
		if (n <= 1) {
			return n;
		}
		int pre2 = 1;
		int pre1 = 1;
		int cur = 0;
		int i = 2;
		while (i <= n) {
			cur = pre1 + pre2;
			int tmp = pre1;
			pre1 = cur;
			pre2 = tmp;
			i++;
		}
		return cur;
	}

}
