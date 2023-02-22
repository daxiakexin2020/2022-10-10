package main

func main() {
	/**
	LISTEN：侦听来自远方TCP端口的连接请求

	SYN-SENT：发送连接请求后等待匹配的连接请求

	SYN-RECEIVED（syn已经收到的）：收到和发送给一个连接请求后等待对连接请求的确认

	ESTABLISHED(已建立的)：代表一个打开的连接，数据可以传送给客户

	FIN-WAIT-1：等待远程TCP的连接中断请求，或者先前的连接中断请求的确认

	FIN-WAIT-2：从远程TCP等待连接中断请求

	todo
	CLOSE-WAIT：等待从本地用户发来的连接中断请求，由被动关闭方持有

	CLOSEING：等待远程TCP对连接中断的确认

	LAST-ACK：等待原来发向远程TCP的连接中断请求的确认

	TIME-WAIT：等待足够的时间一确保远程TCP接受到连接中断请求的确认 2sml 一般是2分钟,由主动关闭方持有

	CLOSED：没有任何连接状态
	*/
}
