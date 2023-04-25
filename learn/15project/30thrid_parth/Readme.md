


分层
    server
        处理业务，例如参数校验
    model
        存放对接客户端的具体的结构体
        将proxy吐出来的结构体，转换，处理
    proxy
        jiuqi
            base.go
                属性 
                    url
                方法
                    send
                    marshal
                    unmarshal
                    
 抽象
    发送
    通用参数组织，例如token，page
    url
    path部分
    返回值 json, 解码在相应的结构体中 ,结构体转化成相应的proto，吐出去
    
