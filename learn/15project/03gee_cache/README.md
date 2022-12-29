分布式缓存

单机缓存和基于 HTTP 的分布式缓存
最近最少访问(Least Recently Used, LRU) 缓存策略
    哈希表+双向链表
使用 Go 锁机制防止缓存击穿
使用一致性哈希选择节点，实现负载均衡
使用 protobuf 优化节点间二进制通信

geecache/
    |--lru/
        |--lru.go  // lru 缓存淘汰策略
    |--byteview.go // 缓存值的抽象与封装
    |--cache.go    // 并发控制
    |--geecache.go // 负责与外部交互，控制缓存存储和获取的主流程