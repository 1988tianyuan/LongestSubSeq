package main

import "fmt"

//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
// 示例:
//
// Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");
//trie.search("app");     // 返回 true
//
// 说明:
//
//
// 你可以假设所有的输入都是由小写字母 a-z 构成的。
// 保证所有输入均为非空字符串。
//
// Related Topics 设计 字典树
// 👍 485 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	subs [26]*Trie
	wordEnd bool		// 标记是否是某个单词的结尾
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	if len(word) == 0 {
		return
	}
	current := this
	position := 0
	for position < len(word) {
		c := word[position]
		if current.subs[c - 'a'] == nil {
			current.subs[c - 'a'] = &Trie{}
		}
		current = current.subs[c - 'a']
		position++
	}
	current.wordEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.findWord(word, true)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.findWord(prefix, false)
}

func (this *Trie) findWord(word string, checkEnd bool) bool {
	if len(word) == 0 {
		return false
	}
	position := 0
	current := this
	for position < len(word) {
		c := word[position]
		node := current.subs[c - 'a']
		if node != nil {
			current = node
			position++
		} else {
			return false
		}
	}
	if checkEnd {
		return current.wordEnd
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.StartsWith("app"))
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	trie.Insert("app")
	fmt.Println(trie.StartsWith("app"))
	fmt.Println(trie.Search("app"))
	trie.Insert("jam")
	fmt.Println(trie.StartsWith("jam"))
	fmt.Println(trie.Search("jam"))
	fmt.Println(trie.Search("am"))
}
