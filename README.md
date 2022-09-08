###使用方法
####说明
    基于gin框架整合的项目骨架 
    使用 wire完成依赖注入
    借鉴kratos框架分层设计

####一、安装

    go install github.com/go-sven/sven/cmd/sven@latest


####二、新建项目

    sven new sven-demo

####三、新增模块

    sven new app/server/order --nomod  
    
    新增order模块，具体开发参照demo 模块 已实现 常规的 数据库操作 事务操作 redis操作等

###分层设计说明
    handler :    定义接口
    service :    DTO层   协同各类 useCase 交互，但是不应处理复杂逻辑
    useCase :    DO层    业务逻辑的组装层 整合 repository cache
    repository : PO层 业务数据访问 mysql redis
    cache      : 处理缓存逻辑

###开发流程
    1.新增handler 文件 比如 handler/index模块下 order.go
        1.1 定义OrderService 接口
        1.2 申明handler的func
    2.新增对应的路由以及需要的中间件

    3.在app/server/order/service 下定义orderService结构体 实现handler 中定义的接口
    
    4.完善 usecase repository model cache等逻辑

    5.完善 wire.go 加入新增的 ProviderSet 执行 [wire] 命令 生成wire_gen.go 文件
    
    6.执行[go run main.go] 测试