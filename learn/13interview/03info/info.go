package main

func main() {
	/**
		mysql:
		     mysql 事务   原子性 一致性 持久性 隔离性
		     mysql索引  B+树
		     执行流程   连接器  解析器  预处理器 优化器 执行器（直接对接引擎层）
	         explain
					id：表示查询中每个操作的序号。
					select_type：表示查询的类型，例如简单查询、联合查询等。
					table：表示相关的表名。
					partitions：表示分区数量，如果涉及分区表的话。
					type：表示访问类型，例如系统、常量、单表查询、联合查询等。
					possible_keys：表示可能使用的索引。
					key：表示实际使用的索引。
					key_len：表示使用的索引长度。
					ref：表示与索引比较的列或常量。
					rows：表示MySQL估计需要扫描的行数。
					filtered：表示按照WHERE条件过滤后的行数占比。
					Extra：表示其他附加信息，例如是否使用了临时表、是否使用了文件排序等。
			 mvcc 保证隔离性   间隙锁
			 引擎类型
					innodb 行锁 支持事务 聚集索引  2个文件 索引和数据在一起
					Myisam 表锁 不支持事务 费聚集索引  3个文件，索引和数据分开的


			golang
				GMP
				channel
				map
				内存溢出  资源忘记回收 gorountine泄露
				垃圾回收  三色法+混合写屏障
				内存逃逸  接口 、 函数返回值*

			redis
				跳跃表
				16384个槽
				hashTable扩容过程

			tcp
				3次握手
				4次挥手
				time_wait(主动关闭方)
				closw_wait（被动关闭方）

			linux
				机器负载突然过高，需要如何排查
								查看占比最高的进程

	*/
}
