package week07

//实现字典树
type Trie struct {
	IsEnd bool
	Next  []*Trie //下一层
}

/** Initialize your data structure here. */
func TrieConstructor() Trie {
	return Trie{
		IsEnd: false,
		Next:  make([]*Trie, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this
	for i := 0; i < len(word); i++ {
		if curr.Next[word[i]-'a'] == nil {
			curr.Next[word[i]-'a'] = &Trie{
				IsEnd: false,
				Next:  make([]*Trie, 26),
			}
		}
		curr = curr.Next[word[i]-'a']
	}
	curr.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curr := this
	for i := 0; i < len(word); i++ {
		if curr.Next[word[i]-'a'] == nil {
			return false
		}
		curr = curr.Next[word[i]-'a']
	}
	if curr.IsEnd == false {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for i := 0; i < len(prefix); i++ {
		if curr.Next[prefix[i]-'a'] == nil {
			return false
		}
		curr = curr.Next[prefix[i]-'a']
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
