http://c.biancheng.net/redis_command/

key 过期
    定时删除
        伴随着每一个key，启动一个定时任务
    定期删除
        固定间隔时间，统一轮询删除
    惰性删除
        访问时，判断key是否过期

库选择 0-16 


key-val
string-
        string int bool
            map[string]string,int,bool
        []set
            map[string][]int,string,bool
        []zset
        list
        hash field value
        

实现
    string  key value
        
    list   list 
            node   {
                key string
                value interface
                next *node
                prev *node
            }
            list{
               head *node
               tail *node 
               data map[string]*node
            }
            lpush rpush lpop rpop 
            
    set key []interface
    sort set key 
    hash key field value
    bitmap map[byte]byte


    cmd
    server
        construct
            string
            set
            zset
            list
            hash
            bitmap
        string
        set
        zset
        list
        hash
        bitmap