# 进行统一的管理，包括认证、转发等等

# 应该具备功能
## 支持tcp、http、https、websocket、grpc等协议
        http：路由、认证、是否应该简单组织业务？？ 、转发    gin
        grpc：
## 网关可以支持多种负载均衡策略：轮询、权重轮询、hash一致性轮询 `
## 支持下游服务发现：主动探测、自动服务发现、注册
    etcd 
## 支持横向扩容：加机器就能解决高并发	
## 公共权限认证
    自己设计
## 限流 
## 熔断
## 降级 


# 模块设计
+ 服务注册
  +  服务名称 string Y
     服务地址 array  Y 
     服务端口 int    Y
     服务类型 string （http｜grpc）  Y
     请求方式 string  (GET POST PUT DELETE OPTIONS HEAD ANY) Y
+ 服务发现
+ 权限验证
  + token验证
  + ip合法性：黑白名单
+ 代理模块，主要负责代理转发请求，比如转发http服务、grpc服务
+ 公共组件，比如注册组件（例如etcd、consol...）
+ 处理模块，主要包含http请求方式、grpc请求方式、tcp请求方式
  + http:
    + 路由：login接口、刷新token接口、请求服务接口...
    + 请求微服务的方式，http
    + 解析请求，path，params，透传
  + grpc
    + 路由：login接口、刷新token接口、请求各个服务接口...
    + 请求微服务的方式，grpc
  + tcp //todo
    

