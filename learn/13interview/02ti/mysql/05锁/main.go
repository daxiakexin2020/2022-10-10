package main

func main() {

	/**
	行锁
		锁住一行的数据
		Innodb 支持，默认
	表锁
		表锁，锁住整张表
		Myisaim

	乐观锁
		只读
	悲观锁
		读写都锁

	间隙锁
		解决幻读
		锁住某一段间隙，比如锁住id=1 -> id=100	之间的数据行
	*/

}