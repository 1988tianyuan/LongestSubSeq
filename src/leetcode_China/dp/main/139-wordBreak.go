package main

import "fmt"

//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
// 说明：
//
//
// 拆分时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
//
//
// 示例 1：
//
// 输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//
//
// 示例 2：
//
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//     注意你可以重复使用字典中的单词。
//
//
// 示例 3：
//
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
//
// Related Topics 动态规划
// 👍 792 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 || len(wordDict) == 0 {
		return false
	}
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			if i - len(word) < -1 {
				continue
			}
			if i - len(word) != -1 && !dp[i - len(word)] {
				continue
			}
			// 判断 i-len(word)+1 ~ i这一段是否在wordDict中
			rangeWord := s[i-len(word)+1:i+1]
			for _, tmpWord := range wordDict {
				if tmpWord == rangeWord {
					dp[i] = true
					break
				}
			}
			if dp[i] {
				break
			}
		}
	}
	return dp[len(s)-1]
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	fmt.Println(wordBreak(s, wordDict))

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}
