package main

func main() {
	/**
	索引分类
		聚蔟索引（主键索引）：只有一个聚簇索引，InnoDB使用聚簇索引，索引树上数据行就存在叶子节点，主键+数据行在一起，找到索引就找到了数据行，不用再此回表
			例如 id pro price title count type 通过主键就可以在叶子节点拿到多有的行数据，不用二次回表了
		    简单示意索引树
	 			1 p001 100   todo
				2 p002 200
				3 p003 300
		二级索引（辅助索引）：可以有多个，索引树上叶子节点，存的是索引值+主键值，而非数据行 todo
			例如 商品标号为二级索引，如下存，二级索引树上的叶子节点p001上放的是主键值1，通过p001索引找到id：1，然后再拿着id：1去主键索引树上找数据行
			简单示意索引树
				p001  1   todo
				p002  2
				p003  3
			由此也可以引申出来，覆盖索引的含义
				例如：此sql： select id from product where pro="p001" 那么，在p的二级索引上存的就是二级索引值和主键id，那么直接返回即可，不用二次回表
				例如：此sql： select id，price from product where pro="p001"，那么price字段的值，并不在二级索引树上，因此，拿到主键id，还需要二次回表到主键索引上拿数据行
			联合索引
			例如：title 与 count、type 3个字段形成联合索引 create index（title，count，type），那么title字段就是最左匹配的字段，想要使用联合索引，sql中必须要有此字段
			简单示意索引树
				title1 count1 type1  1 todo
				title2 count2 type2  2
				title3 count3 type3  3
			前缀索引
				前缀索引是指对字符类型字段的前几个字符建立的索引，而不是在整个字段上建立的索引，前缀索引可以建立在字段类型为 char、 varchar、binary、varbinary 的列上。
				使用前缀索引的目的是为了减少索引占用的存储空间，提升查询效率。

	为什么使用B+树
		1、存储相同量级的数据，B+树比B树低，磁盘IO次数更少  log(N)
		2、B+树叶子节点用双向链表串起来，更适合范围查找，例如Hash索引，进行范围查找，那就是灾难级别的

	什么情况下尽量不创建索引
		1、数据量少
		2、经常更新的字段
		3、在where、order by等子句中少用的字段
		4、值比较少的字段，比如：status：0，1 只有2种状态

	分析sql的工具
		explain

	*/
}
