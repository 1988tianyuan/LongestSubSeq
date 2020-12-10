package string

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
// 示例 4:
//
//
//输入: s = ""
//输出: 0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
// 👍 4688 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	length := len(s)
	longest := 1
	charMap := make(map[uint8]int)
	for i := 0; i < length; i++ {
		if i + longest >= length {
			return longest
		}
		charMap[s[i]] = 1
		curLongest := 1
		for j := i+1; j < length; j++ {
			curChar := s[j]
			if charMap[curChar] == 1 {
				// 已经出现过
				break
			}
			// 没出现过
			charMap[curChar] = 1
			curLongest++
		}
		if curLongest > longest {
			longest = curLongest
		}
		charMap = make(map[uint8]int)
	}
	return longest
}
//leetcode submit region end(Prohibit modification and deletion)

