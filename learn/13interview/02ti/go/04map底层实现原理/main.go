package main

import "runtime"

func main() {
	test()
}

const (
	bucketCntBits = 3 //0000 0011   1+2*1
	bucketCnt     = 1 << bucketCntBits
)

func test() {
	runtime.GC()
	//https://www.topgoer.cn/docs/goquestions/goquestions-1cjh2q3c8knbv
	/**
	todo  结构
		type hmap struct {
			  count     int         	   // 当前哈希表中的元素数量，调用 len(map) 时，直接返回此值
			  flags     uint8
	          B         uint8      	   	   // 表示当前哈希表持有的 buckets 数量，但是因为哈希表中桶的数量都是2的倍数，所以该字段会存储对数，也就是 len(buckets) == 2^B
	          noverflow uint16        	   // overflow 的 bucket 近似数
			  hash0     uint32             // 是哈希的种子，它能为哈希函数的结果引入随机性，这个值在创建哈希表时确定，并在调用哈希函数时作为参数传入
			  buckets    unsafe.Pointer    // 指向 buckets 数组，大小为 2^B 如果元素个数为0，就为 nil
			  oldbuckets unsafe.Pointer    // 是哈希在扩容时用于保存之前 buckets 的字段，它的大小是当前 buckets 的一半；
			  nevacuate  uintptr  		   // 指示扩容进度，小于此地址的 buckets 迁移完成
			  extra *mapextra			   // optional fields
	   }
		map+双向链表，类似lru缓存淘汰机制的结构 ，  进来一个key，首先进行hash函数的计算，计算出应该落的桶的位置，每一个桶中，存的是一个链表，
		最好的结果是0-1个，最多2-3个，再多，性能会下降
		装载因子:=元素数量÷桶数量   装载因子越大，性能越差，如果100%，基本是O(n)的时间复杂度

	todo 扩容
		装载因子已经超过 6.5；
		哈希使用了太多溢出桶；

	
	*/
}
