package main

func main() {

}

func test() {
	/**
	G：任务
	M：抽象线程
	P：控制器，调度器

	P中绑定G，M绑定P，去P中拿任务进行执行
		2个重要机制
			P中的G需要进行IO，操作，保证10ms机制，防止其他的G饿死
			当P中没有G时，不会销毁绑定的M，而是会去其他P中抢占G执行，先去全局队列中拿G
	*/
}
