package gee

import "strings"

type node struct {
	pattern  string  //待匹配路由  例如   /api/rule
	part     string  //路由的一部分，例如  rule
	children []*node //子节点，例如：【doc，intro】
	isWild   bool    //是否精准匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
// parts =>[api test :name]  part=>test
func (n *node) matchChildren(part string) []*node {
	var nodes []*node
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

/**
对于路由来说，最重要的当然是注册与匹配了。
开发服务时，注册路由规则，映射handler；访问时，匹配路由规则，查找到对应的handler。因此，Trie 树需要支持节点的插入与查询。
插入功能很简单，递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个，
有一点需要注意，/p/:lang/doc只有在第三层节点，即doc节点，pattern才会设置为/p/:lang/doc。
p和:lang节点的pattern属性皆为空。因此，当匹配结束时，我们可以使用n.pattern == ""来判断路由规则是否匹配成功。
例如，/p/python虽能成功匹配到:lang，但:lang的pattern值为空，因此匹配失败。
查询功能，同样也是递归查询每一层的节点，退出规则是，匹配到了*，匹配失败，或者匹配到了第len(parts)层节点。


todo  parrent: /api/test/:name
      parts: [api test :name]
	  height: 0
*/

func (n *node) insert(pattern string, parts []string, height int) {

	//递归的出口
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	//todo 拿到第一段，根节点，因为，初始化传入的height是0  , [api test :name]   part=>api
	part := parts[height]

	//返回第一个匹配的孩子
	child := n.matchChild(part)

	//如果没有孩子，则是新建一个孩子，将第一部分设置为新孩子的的part
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	//递归进行第二段匹配，height+1后，下次进来，part=parts[1] part=>test
	child.insert(pattern, parts, height+1)
}

/*
*

	todo
		初始化：
		parts	=> [api test :name]
		height => 0
*/
func (n *node) search(parts []string, height int) *node {

	//递归出口，当height==len(parts)时，代表到了最后一层
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	//第一次进来：     parts=>[api test :name]    height:0   part="api
	part := parts[height]

	//返回part部分包含api的节点 []*node
	children := n.matchChildren(part)

	//递归，接着找匹配到第二段的节点.....
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
