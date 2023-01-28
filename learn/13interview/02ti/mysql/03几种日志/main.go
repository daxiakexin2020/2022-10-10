package main

func main() {

}

func test() {
	/**
	undo log
		它记录了需要回滚的日志信息，事务回滚时撤销已经执行成功的sql
		主要用于事务中执行失败，进行回滚，以及在MVCC中对历史版本的查看
		是逻辑日志
		当一条数据需要更新前，首先会将更新前的数据记录在undo log中，当事务成功提交以后，不会立即清除undo log，而是将其放入待清理列表中

	redo log
		是重做日志文件，记录修改之后的值，用于持久化到磁盘中
		不管事务是否成功，都会记录，保证了事务的持久性

	big log
		用来记录数据库执行写入性的操作，由Server层进行记录，任何使用存储引擎的mysql数据库都会记录bin log日志
		逻辑日志，简答理解，就是sql语句
		用来进行主从复制和数据恢复

	*/
}
