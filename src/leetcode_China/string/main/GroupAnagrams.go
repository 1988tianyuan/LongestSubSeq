package main

import "fmt"

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串
// 👍 608 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	var result [][]string
	tmpMap := make(map[[26]int32][]string, 0)
	for _, str := range strs {
		var letters [26]int32
		for _, l := range str {
			letters[l - 'a']++
		}
		tmpMap[letters] = append(tmpMap[letters], str)
	}
	for _, value := range tmpMap {
		result = append(result, value)
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	array := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(array))
}
