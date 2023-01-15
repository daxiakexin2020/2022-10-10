某个任务的耗时、内存统计组件

type Server struct{
TimeInfo map[string][]time.Time
MemoryInfo map[string][]uint64
}

Add(key string)
//做任务
Print(key string)

Add(key)
End(key)
Print(key)

end()
print()

nowTime() time.Time
nowMemory() uint64


