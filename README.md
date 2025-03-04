# mall
这是一个不完整的商城秒杀系统示例，实用 go-zero,
使用 
* etcd作为服务发现，
* mariadb作为存储，
* redis作为数据缓存和库存持久化，
* kafka作为消息队列，
* dtm 作为分布式事务管理，

主链路
```
seckill(rpc)->
    product(rpc)  检索库存(TODO:这里存在性能瓶颈，可能会击穿到 db)
    kafka 异步消息

seckill(rmq<-kafka)-> 
    order(rpc) 生成订单
    product(rpc) 扣减库存
```