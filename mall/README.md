# mall
## 目录结构
```
//结构
app - BFF服务
cart - 购物车服务
order - 订单服务
pay - 支付服务
product - 商品服务
recommend - 推荐服务
reply - 评论服务
user - 账号服务
在每个服务目录下我们又会分为多个服务，主要会有如下几类服务：
//职责
api - 对外的BFF服务，接受来自客户端的请求，暴露HTTP接口
rpc - 对内的微服务，仅接受来自内部其他微服务或者BFF的请求，暴露gRPC接口
rmq - 负责进行流式任务处理，上游一般依赖消息队列，比如kafka等
admin - 也是对内的服务，区别于rpc，更多的是面向运营侧的且数据权限较高，通过隔离可带来更好的代码级别的安全，直接提供HTTP接口
```
## 初始化
```
mkdir mall 
cd mall
go mod init mall 
```
后续 goctl 命令会用到这个模块名生成代码会使用到 go.mod 里的 module名
```shell
mkdir apps
mkdir apps/order
goctl rpc new rpc  #对内提供 rpc服务
goctl api new admin #对内 http服务
```
