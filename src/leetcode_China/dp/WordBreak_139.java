package leetcode_China.dp;

import java.util.Arrays;
import java.util.List;
import java.util.Objects;

public class WordBreak_139 {

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.wordBreak("applepenapple", Arrays.asList("apple", "pen")));
        System.out.println(solution.wordBreak("catsandog", Arrays.asList("cats", "dog", "sand", "and", "cat")));
    }

    public static class Solution {
        public boolean wordBreak(String s, List<String> wordDict) {
            if (s == null || s.length() == 0 || wordDict == null || wordDict.size() == 0) {
                return false;
            }
            boolean[] dp = new boolean[s.length()];
            for (int i = 0; i < s.length(); i++) {
                for (String word : wordDict) {
                    if (i - word.length() < -1) {
                        continue;
                    }
                    if (i - word.length() != -1 && !dp[i - word.length()]) {
                        continue;
                    }
                    // 判断 i-len(word)+1 ~ i这一段是否在wordDict中
                    String rangeWord = s.substring(i-word.length()+1, i+1);
                    for (String tmpWord : wordDict) {
                        if (Objects.equals(tmpWord, rangeWord)) {
                            dp[i] = true;
                            break;
                        }
                    }
                    if (dp[i]) {
                        break;
                    }
                }
            }
            return dp[s.length()-1];
        }
    }
}
