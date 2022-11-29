package main

func main() {
	test()
}

func test() {

	//https://www.topgoer.cn/docs/goquestions/goquestions-1cjh2q3c8knbv

	/**
		type hmap struct {

	     count     int         // 元素个数，调用 len(map) 时，直接返回此值

		  flags     uint8

	       B         uint8         // buckets 的对数 log_2

	       noverflow uint16                // overflow 的 bucket 近似数

	       hash0     uint32        // 计算 key 的哈希的时候会传入哈希函数

	          // 指向 buckets 数组，大小为 2^B
	       // 如果元素个数为0，就为 nil
	       buckets    unsafe.Pointer

	       // 扩容的时候，buckets 长度会是 oldbuckets 的两倍
	       oldbuckets unsafe.Pointer

	       // 指示扩容进度，小于此地址的 buckets 迁移完成
	       nevacuate  uintptr

	       extra *mapextra // optional fields
	   }
	*/
}
