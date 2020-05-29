package leetcode_China.dp;

import java.util.HashMap;

/**
 * 字符串 s ，请你返回满足以下条件的最长子字符串的长度：
 * 	 每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次
 * 	 
 * 输入：s = "eleetminicoworoep"
 * "ebvobnleetmnbne"
 * 输出：13
 * 解释：最长子字符串是 "leetminicowor" ，它包含 e，i，o 各 2 个，以及 0 个 a，u 。
 */
public class LongestSubString {
	
	public static int findTheLongestSubstring(String s) {
		int len = s.length();
		Integer[] locations = new Integer[0b11111 + 1];
		int longest = 0;
		int current = 0;
		for (int i = 0; i < len; i++) {
			char currentChar = s.charAt(i);
			switch (currentChar) {
				case 'a':
					current ^= 0b00001;
					break;
				case 'e':
					current ^= 0b00010;
					break;
				case 'i':
					current ^= 0b00100;
					break;
				case 'o':
					current ^= 0b01000;
					break;
				case 'u':
					current ^= 0b10000;
					break;
			}
			if (locations[current] == null) {
				locations[current] = i;
			}
			if (current == 0) {
				longest = i + 1;
			} else {
				int exists;
				if ((exists = locations[current]) != i) {
					longest = Math.max(i - exists, longest);
				}
			}
		}
		return longest;
	}
	
	public static void main(String[] args) {
		System.out.println(findTheLongestSubstring("eleetminicoworoep"));
	}
}
