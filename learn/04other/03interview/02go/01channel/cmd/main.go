package main

func main() {

}

func HandleChannel() {
	/**
	TODO
		https://blog.csdn.net/h_l_f/article/details/118255931
		type hchan struct {
		qcount   uint   // channel 里的元素计数
		dataqsiz uint   // 可以缓冲的数量，如 ch := make(chan int, 10)。 此处的 10 即 dataqsiz
		elemsize uint16 // 要发送或接收的数据类型大小
		buf      unsafe.Pointer // 当 channel 设置了缓冲数量时，该 buf 指向一个存储缓冲数据的区域，该区域是一个循环队列的数据结构
		closed   uint32 // 关闭状态
		sendx    uint  // 当 channel 设置了缓冲数量时，数据区域即循环队列此时已发送数据的索引位置
		recvx    uint  // 当 channel 设置了缓冲数量时，数据区域即循环队列此时已接收数据的索引位置
		recvq    waitq // 想读取数据但又被阻塞住的 goroutine 队列
		sendq    waitq // 想发送数据但又被阻塞住的 goroutine 队列
		lock mutex
		...
	}
	*/
}
