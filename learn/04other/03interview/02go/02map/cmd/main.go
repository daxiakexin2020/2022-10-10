package main

func main() {

}

func HandleMap() {

	/*
		TODO
			https://blog.csdn.net/u010177891/article/details/105985906
			map无序的key-value结构
			type hmap struct {
			// map中元素的个数，使用len返回就是该值
			count     int
			// 状态标记
			// 1: 迭代器正在操作buckets
			// 2: 迭代器正在操作oldbuckets
			// 4: go协程正在像map中写操作
			// 8: 当前的map正在增长，并且增长的大小和原来一样
			flags     uint8
			// buckets桶的个数为2^B
			B         uint8
			// 溢出桶的个数
			noverflow uint16
			// key计算hash时的hash种子
			hash0     uint32
			// 指向的是桶的地址
			buckets    unsafe.Pointer
			// 旧桶的地址，当map处于扩容时旧桶才不为nil
			oldbuckets unsafe.Pointer
			// map扩容时会逐步讲旧桶的数据迁移到新桶中，此字段记录了旧桶中元素的迁移个数当 nevacuate>=旧桶元素个数时数据迁移完成
			nevacuate  uintptr
			// 扩展字段
			extra *mapextra
		}
	*/
}
