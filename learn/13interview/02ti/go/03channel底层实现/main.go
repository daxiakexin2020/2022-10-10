package main

func main() {

	/**
	todo	设计思想
		不要通过共享内存，进行通信，而是通过通信，共享内存
		1.传统的模式：共享内存通信  Thread1=》 内存 《=Thread2
		2.Go的CSP模式：通过channel进行通信   Goroutine1=》 Channel 《=Goroutine2

	todo	Channel通信原则：先入先出  FIFO
			先从channel读取数据的Goroutine会先接收到数据
			先向channle发送数据的Goroutine会先得到发送数据的权利

	todo	同步 Channel — 不需要缓冲区，发送方会直接将数据交给（Handoff）接收方；
			异步 Channel — 基于环形缓存的传统生产者消费者模型；
			chan struct{} 类型的异步 Channel — struct{} 类型不占用内存空间，不需要实现缓冲区和直接发送（Handoff）的语义；

	todo	数据结构
				type hchan struct {
					qcount   uint   			channel中元素的个数
					dataqsiz uint   			channel循环队列的长度
					buf      unsafe.Pointer 	channel缓冲区数据指针
					elemsize uint16    			当前 Channel 能够收发的元素大小
					closed   uint32				channel是否关闭
					elemtype *_type				当前 Channel 能够收发的元素类型
					sendx    uint    			channel发送操作处理到的位置
					recvx    uint	 			channel接收操作处理到的位置
					recvq    waitq				接收数据队列		（双向链表） 当前 Channel 由于缓冲区空间不足而阻塞的 Goroutine 列表
					sendq    waitq				发送数据队列		（双向链表） 当前 Channel 由于缓冲区空间不足而阻塞的 Goroutine 列表
					lock mutex					互斥锁
												双向链表  sudog 表示一个在等待列表中的 Goroutine，该结构中存储了两个分别指向前后 runtime.sudog 的指针以构成链表。
												type waitq struct {
													first *sudog
													last  *sudog
												}
				}

	todo	流程
				1、创建管道
				Go 语言中所有 Channel 的创建都会使用 make 关键字。
				编译器会将 make(chan int, 10) 表达式转换成 OMAKE 类型的节点，并在类型检查阶段将 OMAKE 类型的节点转换成 OMAKECHAN 类型：
				func makechan(t *chantype, size int) *hchan {
						elem := t.elem
						mem, _ := math.MulUintptr(elem.size, uintptr(size))
						var c *hchan
						switch {
						case mem == 0:
							c = (*hchan)(mallocgc(hchanSize, nil, true))
							c.buf = c.raceaddr()
						case elem.kind&kindNoPointers != 0:
							c = (*hchan)(mallocgc(hchanSize+mem, nil, true))
							c.buf = add(unsafe.Pointer(c), hchanSize)
						default:
							c = new(hchan)
							c.buf = mallocgc(mem, elem, true)
						}
						c.elemsize = uint16(elem.size)
						c.elemtype = elem
						c.dataqsiz = uint(size)
						return c
				}
			2、发送数据

	*/

}
