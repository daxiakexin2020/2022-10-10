微服务
服务注册与发现
https://blog.csdn.net/energysober/article/details/124026804
https://blog.csdn.net/kankan231/article/details/126083953
https://blog.csdn.net/kankan231/article/details/126212654
https://docs.microsoft.com/zh-cn/dotnet/architecture/grpc-for-wcf-developers/protobuf-data-types
https://blog.csdn.net/kevin_tech/article/details/109281835
https://blog.csdn.net/u013862108/article/details/103808804
https://github.com/goeasya/discox
Etcd
https://blog.csdn.net/m0_58541541/article/details/123183979

Protoc
安装：https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3
帖子：
https://blog.csdn.net/u012830303/article/details/126517355
https://www.cnblogs.com/zhangmingcheng/p/16329237.html
https://blog.csdn.net/wxystyle/article/details/124297960
protoc-gen-go安装:
https://www.cnblogs.com/ginkgo-leaf/p/16223111.html
go install github.com/golang/protobuf/protoc-gen-go@latest

protoc-gen-go-grpc安装：
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

demo:
//两个个文件
protoc --go_out=.  message.proto
protoc --go-grpc_out=. message.proto

//一个文件
protoc --go_out=plugins=grpc:. user_grpc.proto
