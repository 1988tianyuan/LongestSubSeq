package main

import "fmt"

//给定两个字符串 s 和 t，判断它们是否是同构的。
//
// 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
//
// 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
//
// 示例 1:
//
// 输入: s = "egg", t = "add"
//输出: true
//
//
// 示例 2:
//
// 输入: s = "foo", t = "bar"
//输出: false
//
// 示例 3:
//
// 输入: s = "paper", t = "title"
//输出: true
//
// 说明:
//你可以假设 s 和 t 具有相同的长度。
// Related Topics 哈希表
// 👍 295 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isIsomorphic(s string, t string) bool {
	aMap := make(map[byte]byte, 0)
	for i := 0; i < len(s); i++ {
		sA := s[i]
		tA := t[i]
		if aMap[sA] == 0 {
			aMap[sA] = tA
		} else {
			if aMap[sA] != tA {
				return false
			}
		}
	}
	bMap := make(map[byte]byte, 0)
	for i := 0; i < len(s); i++ {
		sA := s[i]
		tA := t[i]
		if bMap[tA] == 0 {
			bMap[tA] = sA
		} else {
			if bMap[tA] != sA {
				return false
			}
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	s := "paper"
	t := "title"
	fmt.Println(isIsomorphic(s, t))
	s2 := "foo"
	t2 := "bar"
	fmt.Println(isIsomorphic(s2, t2))
	s3 := "egg"
	t3 := "add"
	fmt.Println(isIsomorphic(s3, t3))
	s4 := "ab"
	t4 := "aa"
	fmt.Println(isIsomorphic(s4, t4))
}
