/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	isEnd bool
	Next  [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	now := this

	for _, ch := range word {
		if now.Next[ch-'a'] == nil {
			now.Next[ch-'a'] = &Trie{}
		}

		now = now.Next[ch-'a']
	}

	now.isEnd = true
}

func (this *Trie) Search(word string) bool {
	now := this
	for _, ch := range word {
		if now.Next[ch-'a'] == nil {
			return false
		}

		now = now.Next[ch-'a']
	}

	return now.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	now := this
	for _, ch := range prefix {
		if now.Next[ch-'a'] == nil {
			return false
		}

		now = now.Next[ch-'a']
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
