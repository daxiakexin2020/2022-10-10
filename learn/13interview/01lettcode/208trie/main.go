package main

import "fmt"

func main() {
	test()
}

func test() {
	t := Constructor()
	t.Insert("abc")
	t.Insert("def")
	res := t.Search("ab")
	fmt.Printf("res:=%v", res)
}

/**
方法一：字典树
\text{Trie}Trie，又称前缀树或字典树，是一棵有根树，其每个节点包含以下字段：

指向子节点的指针数组 \textit{children}children。对于本题而言，数组长度为 2626，即小写英文字母的数量。
此时 \textit{children}[0]children[0] 对应小写字母 aa，\textit{children}[1]children[1] 对应小写字母 bb，…，\textit{children}[25]children[25] 对应小写字母 zz。
布尔字段 \textit{isEnd}isEnd，表示该节点是否为字符串的结尾。

todo 插入字符串

我们从字典树的根开始，插入字符串。对于当前字符对应的子节点，有两种情况：

子节点存在。沿着指针移动到子节点，继续处理下一个字符。
子节点不存在。创建一个新的子节点，记录在 \textit{children}children 数组的对应位置上，然后沿着指针移动到子节点，继续搜索下一个字符。
重复以上步骤，直到处理字符串的最后一个字符，然后将当前节点标记为字符串的结尾。

todo 查找前缀

我们从字典树的根开始，查找前缀。对于当前字符对应的子节点，有两种情况：

子节点存在。沿着指针移动到子节点，继续搜索下一个字符。
子节点不存在。说明字典树中不包含该前缀，返回空指针。
重复以上步骤，直到返回空指针或搜索完前缀的最后一个字符。

若搜索到了前缀的末尾，就说明字典树中存在该前缀。此外，若前缀末尾对应节点的 \textit{isEnd}isEnd 为真，则说明字典树中存在该字符串。


*/

type Trie struct {
	children [26]*Trie //指向子节点的指针数组 \textit{children}children。对于本题而言，数组长度为 26，children[0] 对应小写字母 a,children[1] 对应小写字母 b，children[25] 对应小写字母 z
	isEnd    bool      //布尔字段 isEnd，表示该节点是否为字符串的结尾。
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word { //abc   a,b,c    97 ， 98，99
		ch -= 'a'                     //0   1   2
		if node.children[ch] == nil { //不存在0这个节点，添加一个节点
			node.children[ch] = &Trie{}
		}
		node = node.children[ch] //移动node，指向刚刚新添加的节点
	}

	//t：a=》b=》c

	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix { //a
		ch -= 'a' //0
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
