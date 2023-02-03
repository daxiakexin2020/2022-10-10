package main

func main() {
	/**
	使用bit位标志某个值是否存在
	redis使用bitmap实现
	实现原理：
		key1->hash1(key1)->hashCode1
		key1->hash2(key1)->hashCode2
		key1->hash3(key1)->hashCode3
		使用几个hash函数，例如使用3个hash函数，计算code，根据哈希值计算出一个整数索引值，将该索引值与位数组长度做取余运算，最终得到一个位数组位置，并将该位置的值变为 1
		每个 hash 函数都会计算出一个不同的位置，然后把数组中与之对应的位置变为 1。通过上述过程就完成了元素添加(add)操作。

	存在问题：
		元素可能存在，因为存在hash冲突的可能，
	*/
}
