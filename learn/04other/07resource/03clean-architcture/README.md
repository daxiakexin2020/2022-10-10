此代码对应文章：Golang 简洁架构实战 https://www.luozhiyun.com/archives/640


api(服务，最外层)，注入的是services，接口的形式
    services注入的是repo(数据源)，接口形式
        repo是对接各种数据源，注入的是例如DB，Redis，这里可以再封装一层，注入数据源，接口形式