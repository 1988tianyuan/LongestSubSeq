package leetcode_China.string;

/**
 * 写一个函数 StrToInt，实现把字符串转换成整数这个功能
 */
public class StringToInteger {
	
	public static int strToInt(String str) {
		if (str == null || str.length() == 0) {
			return 0;
		}
		char[] chars = str.toCharArray();
		int beginIndex = -1;
		int size = chars.length;
		boolean isPositive = true;
		boolean hasNum = false;
		long result = 0;
		for (int i = 0; i < size; i++) {
			char c = chars[i];
			if (Character.isDigit(c)) {
				if (beginIndex == -1) {
					beginIndex = i;
					hasNum = true;
				}
				result = result * 10 + Character.getNumericValue(c);
				if (result > Integer.MAX_VALUE) {
					return isPositive ? Integer.MAX_VALUE : Integer.MIN_VALUE;
				}
			} else {
				if (hasNum) {
					break;
				}
				if (c != ' ') {
					if (c == '-' || c == '+') {
						if (i != size-1 && !Character.isDigit(chars[i+1])) {
							return 0;
						}
						if (c == '-') {
							isPositive = false;
						}
					} else {
						return 0;
					}
				}
			}
		}
		if (!isPositive) {
			return (int) (0 - result);
		}
		return (int)result;
	}
	
	public static void main(String[] args) {
		System.out.println(strToInt("   -34  dasd  "));
	}
}
