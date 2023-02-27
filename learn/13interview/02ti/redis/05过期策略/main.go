package main

func main() {
	/**
	1、定期删除
		默认100ms就定期检查一批过期的key，过期就删除，这里是随机进行检查的
	2、惰性删除
		在你获取某个key的时候，redis会检查一下，这个key是否过期，如果设置了过期时间，并且已经过期，就将其删除了
			对内存不友好
	3、定时删除
		在设定key的过期时间时，同时设置一个定时器，当过期时间来临时，进行删除
			对CPU不友好
	redis采用惰性删除+定期删除，这个组合进行过期策略
	*/
}
