package main

func main() {

	/**
	todo 设计思想
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
												type sudog struct {
													next *sudog
													prev *sudog
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
				当我们需要向channel发送数据时，就会使用类似这样的操作   ch<-i ,编译器会将其解析成OSEND节点，操作过程中，比如会检查管道是否关闭
				func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
					1. 加锁
						lock(&c.lock)
					2. 判断管道是否关闭
						if c.closed != 0 {
							unlock(&c.lock)
							panic(plainError("send on closed channel"))
						}
					3. 如果目标 Channel 没有被关闭并且已经有处于读等待的 Goroutine，那么 runtime.chansend 会从接收队列 recvq 中取出最先陷入等待的 Goroutine 并直接向它发送数据
						if sg := c.recvq.dequeue(); sg != nil {
							send(c, sg, ep, func() { unlock(&c.lock) }, 3)
							return true
						}
					4.缓冲区，如果创建的 Channel 包含缓冲区并且 Channel 中的数据没有装满，会执行下面流程
						在这里我们首先会使用 runtime.chanbuf 计算出下一个可以存储数据的位置，然后通过 runtime.typedmemmove 将发送的数据拷贝到缓冲区中并增加 sendx 索引和 qcount 计数器。
						if c.qcount < c.dataqsiz {
								qp := chanbuf(c, c.sendx)
								typedmemmove(c.elemtype, qp, ep)
								c.sendx++
								if c.sendx == c.dataqsiz {
									c.sendx = 0
								}
								c.qcount++
								unlock(&c.lock)
								return true
						}
				}
			3、接收数据
				直接接收 #
					当 Channel 的 sendq 队列中包含处于等待状态的 Goroutine 时，该函数会取出队列头等待的 Goroutine，处理的逻辑和发送时相差无几，
					只是发送数据时调用的是 runtime.send 函数，而接收数据时使用 runtime.recv
			4、关闭管道
				编译器会将用于关闭管道的 close 关键字转换成 OCLOSE 节点以及 runtime.closechan 函数。
				func closechan(c *hchan) {
					if c == nil {
						panic(plainError("close of nil channel")) 当 Channel 是一个空指针或者已经被关闭时，Go 语言运行时都会直接崩溃并抛出异常
					}
					lock(&c.lock)
					if c.closed != 0 {
						unlock(&c.lock)
						panic(plainError("close of closed channel"))
					}
				}
	*/

}
