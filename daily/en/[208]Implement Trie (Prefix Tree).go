//Trie (we pronounce "try") or prefix tree is a tree data structure used to retr
//ieve a key in a strings dataset. There are various applications of this very eff
//icient data structure, such as autocomplete and spellchecker.
//
// Implement the Trie class:
//
//
// Trie() initializes the trie object.
// void insert(String word) inserts the string word to the trie.
// boolean search(String word) returns true if the string word is in the trie (i
//.e., was inserted before), and false otherwise.
// boolean startsWith(String prefix) returns true if there is a previously inser
//ted string word that has the prefix prefix, and false otherwise.
//
//
//
// Example 1:
//
//
//Input
//["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
//[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
//Output
//[null, null, true, false, true, null, true]
//
//Explanation
//Trie trie = new Trie();
//trie.insert("apple");
//trie.search("apple");   // return True
//trie.search("app");     // return False
//trie.startsWith("app"); // return True
//trie.insert("app");
//trie.search("app");     // return True
//
//
//
// Constraints:
//
//
// 1 <= word.length, prefix.length <= 2000
// word and prefix consist of lowercase English letters.
// At most 3 * 104 calls will be made to insert, search, and startsWith.
//
// Related Topics Design Trie
// 👍 4323 👎 68

package en

//leetcode submit region begin(Prohibit modification and deletion)
const R = 26

type Trie struct {
	//value int
	isEnd bool
	next  [R]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{next: [R]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		value := int(c - 'a')
		nextTrie := node.next[value]
		if nextTrie == nil {
			nextTrie = &Trie{next: [R]*Trie{}}
			node.next[value] = nextTrie
		}
		node = nextTrie
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		nextTrie := node.next[int(c-'a')]
		if nextTrie == nil {
			return false
		}
		node = nextTrie
	}

	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		nextTrie := node.next[int(c-'a')]
		if nextTrie == nil {
			return false
		}
		node = nextTrie
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
//leetcode submit region end(Prohibit modification and deletion)
