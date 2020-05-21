package leetcode_China;

/**
 * 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
 */
public class SumNums {
	
	public static void main(String[] args) {
		System.out.println(sumNums(3));
	}
	
	public static int sumNums(int n) {
		int upper = (int)Math.pow(n + 1, 2) >> 1;
		int lower = (n+1) >> 1;
		return upper - lower;
	}
}
