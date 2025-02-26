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
## 添加 apps/app api
bff接口对外 提供 http api

mall/apps/app/api/api.api

```
cd mall/apps/app/api
goctl api go -api api.api -dir .
```

## 添加  order/rpc order.proto
```
rm -rf mall/apps/order/rpc
mall/apps/order/rpc/order.proto
goctl rpc protoc order.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```
运行  需先启动 etcd
```shell
go run order.go
```
输出
```shell
Starting rpc server at 0.0.0.0:8080...
{"@timestamp":"2025-02-22T09:48:53.173+08:00","caller":"stat/usage.go:61","content":"CPU: 0m, MEMORY: Alloc=3.2Mi, TotalAlloc=5.2Mi, Sys=14.3Mi, NumGC=3","level":"stat"}
{"@timestamp":"2025-02-22T09:48:53.186+08:00","caller":"load/sheddingstat.go:61","content":"(rpc) shedding_stat [1m], cpu: 0, total: 0, pass: 0, drop: 0","level":"stat"}
{"@timestamp":"2025-02-22T09:49:53.174+08:00","caller":"stat/usage.go:61","content":"CPU: 0m, MEMORY: Alloc=3.4Mi, TotalAlloc=5.3Mi, Sys=14.3Mi, NumGC=3","level":"stat"}
{"@timestamp":"2025-02-22T09:49:53.185+08:00","caller":"load/sheddingstat.go:61","content":"(rpc) shedding_stat [1m], cpu: 0, total: 0, pass: 0, drop: 0","level":"stat"}
```
添加业务逻辑
/mall/apps/order/rpc/internal/logic/orderlogic.go
## 添加  product/rpc product.proto
步骤和上面类似，
需要注意修改 mall/apps/product/rpc/etc  下的端口和 order不同即可，避免端口冲突

添加业务逻辑
/mall/apps/product/rpc/internal/logic/productlogic.go

## 修改 apps/app api
修改其/etc配置，添加 order product rpc 信息

在internal/config、svc 添加 rpcclient相关信息

在 internal/logic 添加逻辑（会调用 rpc）

##  test1 
启动 product order rpc，启动 bff api 

test
```shell
http://127.0.0.1:8888/v1/order/list?uid=123

{
  "orders": [
    {
      "order_id": "20220609123456",
      "status": 0,
      "quantity": 0,
      "payment": 0,
      "total_price": 0,
      "create_time": 0,
      "product_id": 0,
      "product_name": "测试商品名称",
      "product_image": "",
      "product_description": ""
    }
  ],
  "is_end": false,
  "order_time": 0
}
```
## logx
mall/apps/app/api/api.go

设置关闭日志 stat,设置错误级别（取消 info,debug)
```shell
func init() {
//logx.DisableStat()
//logx.SetLevel(logx.ErrorLevel)
}
```

logic 

添加字段
```shell
logx.Infow("order list", logx.Field("uid",req.UID))
```
## 弱依赖
NonBlock: true 不影响服务主服务启动，否则启动失败
```shell
ReplyRPC:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: reply.rpc
  NonBlock: true
```
## db
导入 db, 

生成 model
```shell
goctl model mysql datasource -url="user:pass@tcp(127.0.0.1:3306)/product" --dir="./apps/product/rpc/internal/model" -cache=true -table="product,category"
```
开启redis缓存默认缓存时间 7 天 后如果model层发生 update 操作，会自动删除对应缓存
,如果没有查到数据会设置一个空缓存，空缓存的过期时间为1分钟

查询中可能会在时间格式问题，TIMESTAMP解析为time.Time需
dsn 设置parseTime参数  user:password@tcp(127.0.0.1:3306)/dbname?parseTime=true

product下的
etc svc config 修改相关配置和代码
修改 logic代码

## mr.MapReduce
logic 中使用 mr.MapReduce 并发执行 map reduce 函数，如果 map，reduce过程出错可以调用 cancel 取消当前map,reduce 所有任务

## mr.Finish
并行运行函数，如果有任何错误则取消

## image upload
product admin[rest api]

使用了 aliyun oss

## grpc test
配置里需Mode: dev,
gRPC 服务器才会注册反射服务。如果没有启用反射，grpcurl 将无法直接通过服务端查询服务的定义，除非你手动提供 .proto 文件或 protoset 文件。
```shell
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

```shell
grpcurl -plaintext 127.0.0.1:8081  list product.Product
product.Product.Product
product.Product.ProductList
product.Product.Products

grpcurl -plaintext -d '{"product_id": 1}' 127.0.0.1:8081 product.Product.Product
{
  "productId": "1",
  "name": "name"
}
```
未开启 Mode:dev 时需携带 proto
```shell
grpcurl -proto apps/product/rpc/product.proto -plaintext 127.0.0.1:8081  list product.Product 
```
## 自定义 db 方法
mall/apps/product/rpc/internal/model/productmodel.go

customProductModel 实现了 productModel

1. 添加接口方法
2. 实现方法

## redis
实现分类商品缓存 zset

mall/apps/product/rpc/internal/svc

svc内 :redis.MustNewRedis(c.BizRedis)

## SingleGroup
防止热点数据缓存击穿

mall/apps/product/rpc/internal/logic/productlogic.go

l.svcCtx.SingleGroup.Do

## 缓存穿透
从 cache 穿透到 db

默认带 cache db 查询无数据时 缓存一分钟

## db  熔断[自适应]
github.com/zeromicro/go-zero@v1.8.0/core/stores/sqlx/sqlconn.go:117
```go
func (db *commonSqlConn) ExecCtx(ctx context.Context, q string, args ...any) (
	result sql.Result, err error) {
	ctx, span := startSpan(ctx, "Exec")
	defer func() {
		endSpan(span, err)
	}()

	err = db.brk.DoWithAcceptableCtx(ctx, func() error {
		var conn *sql.DB
		conn, err = db.connProv()
		if err != nil {
			db.onError(ctx, err)
			return err
		}

		result, err = exec(ctx, conn, q, args...)
		return err
	}, db.acceptable)
	if errors.Is(err, breaker.ErrServiceUnavailable) {
		metricReqErr.Inc("Exec", "breaker")
	}

	return
}
```
## localcache
```go
collection.NewCache(localCacheExpire)
```








