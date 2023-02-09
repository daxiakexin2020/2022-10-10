package main

func main() {
	/**
	string
		SDS(简单动态字符串)，区别与C的字符串，redis自封装了一层：伪代码：
		struct{
			len	字符串长度
			alloc	分配的空间的长度
			flags	sds类型
			buf[]	字节数组
		}

	list
		quicklist(快速列表)
		typedef struct listNode {
			//前置节点
			struct listNode *prev;
			//后置节点
			struct listNode *next;
			//节点的值
			void *value;
		} listNode;

	hash
		hash表+listpack
		typedef struct dictht {
			//哈希表数组
			dictEntry **table;
			//哈希表大小
			unsigned long size;
			//哈希表大小掩码，用于计算索引值
			unsigned long sizemask;
			//该哈希表已有的节点数量
			unsigned long used;
		} dictht;
		链式hash解决冲突

	zset
		跳跃表 时间复杂度O(logN) + hash表
		跳跃表是为了高效的进行范围查询
		每个节点的层数，是随机生成的
		todo interview： 为什么用跳表，不用其他平衡树？
			1 它们不是非常内存密集型的。基本上由你决定。改变关于节点具有给定级别数的概率的参数将使其比 btree 占用更少的内存。
			2 Zset 经常需要执行 ZRANGE 或 ZREVRANGE 的命令，即作为链表遍历跳表。通过此操作，跳表的缓存局部性至少与其他类型的平衡树一样好。
			3 它们更易于实现、调试等。例如，由于跳表的简单性，我收到了一个补丁（已经在Redis master中），其中扩展了跳表，在 O(log(N) 中实现了 ZRANK。它只需要对代码进行少量修改。
			4 从内存占用上来比较，跳表比平衡树更灵活一些。平衡树每个节点包含 2 个指针（分别指向左右子树），而跳表每个节点包含的指针数目平均为 1/(1-p)，具体取决于参数 p 的大小。如果像 Redis里的实现一样，取 p=1/4，那么平均每个节点包含 1.33 个指针，比平衡树更有优势。
			5 在做范围查找的时候，跳表比平衡树操作要简单。在平衡树上，我们找到指定范围的小值之后，还需要以中序遍历的顺序继续寻找其它不超过大值的节点。如果不对平衡树进行一定的改造，这里的中序遍历并不容易实现。而在跳表上进行范围查找就非常简单，只需要在找到小值之后，对第 1 层链表进行若干步的遍历就可以实现。
			6 从算法实现难度上来比较，跳表比平衡树要简单得多。平衡树的插入和删除操作可能引发子树的调整，逻辑复杂，而跳表的插入和删除只需要修改相邻节点的指针，操作简单又快速。
	*/
}
