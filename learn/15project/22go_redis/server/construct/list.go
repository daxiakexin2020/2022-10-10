package construct

import "fmt"

type Clist struct {
	head *node
	tail *node
	size int32
	*exTime
}

type node struct {
	val  interface{}
	next *node
	prev *node
}

var _ VInterface = (*Clist)(nil)

func (cl *Clist) Type() string {
	return LIST
}

func NewClist(val interface{}) *Clist {
	n := newNode(val)
	return &Clist{
		head: n,
		tail: n,
		size: 1,
	}
}

func newNode(val interface{}) *node {
	return &node{val: val}
}

func (cl *Clist) GetVal() interface{} {
	return cl.head
}

func (cl *Clist) SetVal(data interface{}) {}

func (cl *Clist) AddToHead(val interface{}) {
	cl.addToHead(newNode(val))
}

func (cl *Clist) AddToTail(val interface{}) {
	cl.addToTail(newNode(val))
}

func (cl *Clist) addToHead(node *node) {
	if cl.head == nil {
		cl.initFromNode(node)
	} else {
		node.next = cl.head
		cl.head.prev = node
		cl.head = node
		cl.size++
	}
}

func (cl *Clist) addToTail(node *node) {
	if cl.tail == nil {
		cl.initFromNode(node)
	} else {
		node.prev = cl.tail
		cl.tail.next = node
		cl.tail = node
		cl.size++
	}
}

func (cl *Clist) initFromNode(node *node) {
	cl.head = node
	cl.tail = node
	cl.size = 1
}

func (cl *Clist) RemoveHead() *node {
	if cl.head == nil {
		return nil
	}
	if cl.size == 0 {
		return nil
	}
	n := cl.head
	if cl.head.next != nil {
		cl.head = cl.head.next
		cl.head.prev = nil
	} else {
		cl.head = nil
		cl.tail = nil
	}
	cl.size--
	return n
}

func (cl *Clist) RemoveTail() *node {
	if cl.tail == nil {
		return nil
	}
	if cl.size == 0 {
		return nil
	}
	node := cl.tail
	if cl.tail.prev != nil {
		cl.tail = cl.tail.prev
		cl.tail.next = nil
	} else {
		cl.tail = nil
		cl.head = nil
	}
	cl.size--
	return node
}

func (cl *Clist) GetHead() *node {
	return cl.head
}

func (cl *Clist) GetTail() *node {
	return cl.tail
}

func (cl *Clist) GetHeadVal() interface{} {
	if cl.head == nil {
		return nil
	}
	return cl.head.val
}

func (cl *Clist) GetTailVal() interface{} {
	if cl.tail == nil {
		return nil
	}
	return cl.tail.val
}

func (cl *Clist) Size() int32 {
	return cl.size
}

func (cl *Clist) Get(index int) *node {
	next := cl.head
	for i := 0; i < index; i++ {
		if next.next != nil {
			next = next.next
		} else {
			return nil
		}
	}
	return next
}

func (cl *Clist) GetRange(start int, end int) []*node {

	var nodes []*node
	next := cl.head
	for i := 0; i < end; i++ {
		if i >= start {
			nodes = append(nodes, next)
		}
		if next.next != nil {
			next = next.next
		} else {
			return nodes
		}
	}
	return nodes
}

func (cl *Clist) Remove(index int) *node {
	n := cl.Get(index)
	if n == nil {
		return nil
	}
	prev := n.prev
	next := n.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	cl.size--
	return n
}

func (cl *Clist) RemoveVal(count int32, val interface{}) int32 {
	if count == 0 {
		count = cl.size
	}

	var rcount int32
	curr := cl.head
	if curr == nil {
		return rcount
	}

	for curr != nil && rcount < count {
		if curr.val == val {
			rcount++
			prev := curr.prev
			next := curr.next
			if prev != nil {
				prev.next = next
			} else {
				cl.head = next
			}
			if next != nil {
				next.prev = prev
			} else {
				cl.tail = prev
			}
			cl.size--
		}
		curr = curr.next
	}
	return rcount
}

func (n *node) GetVal() interface{} {
	return n.val
}

func (n *node) SetVal(val interface{}) {
	n.val = val
}

func (cl *Clist) ShowFromHead() {
	f := cl.head
	for f != nil {
		fmt.Println("val:", f.val)
		f = f.next
	}
	fmt.Println("*****************head over*****************")
}

func (cl *Clist) ShowFromTail() {
	f := cl.tail
	for f != nil {
		fmt.Println("val:", f.val)
		f = f.prev
	}
	fmt.Println("*****************tail over*****************")
}
