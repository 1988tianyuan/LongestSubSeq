package main

import "fmt"

//给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
//
// 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
//
// 说明：
//
//
// 字母异位词指字母相同，但排列不同的字符串。
// 不考虑答案输出的顺序。
//
//
// 示例 1:
//
//
//输入:
//s: "cbaebabacd" p: "abc"
//
//输出:
//[0, 6]
//
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
//
//
// 示例 2:
//
//
//输入:
//s: "abab" p: "ab"
//
//输出:
//[0, 1, 2]
//
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
//
// Related Topics 哈希表
// 👍 443 👎 0

//s: "cbaebabacd" p: "abc"
//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	length := len(p)
	result := make([]int, 0)
	pArray := make([]int, 26)
	tmpArray := make([]int, 26)
	for _,a := range p {
		a = a - 'a'
		pArray[a] = pArray[a] + 1
	}
	for i := 0; i + length <= len(s); i++ {
		if isAnagram(s, i, pArray, tmpArray, len(p)) {
			result = append(result, i)
		}
	}
	return result
}

func isAnagram(s string, start int, pArray []int, tmpArray []int, length int) bool {
	for i := start; i < start + length; i++ {
		a := s[i] - 'a'
		tmpArray[a] = tmpArray[a] + 1
	}
	isAnagram := true
	for i := 0; i < 26; i++ {
		if pArray[i] != tmpArray[i] {
			isAnagram = false
		}
		tmpArray[i] = 0
	}
	return isAnagram
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

