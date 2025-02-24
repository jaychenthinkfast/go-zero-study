api[after,order,product,reply]:
	cd apps/app/api && go run api.go

order:
	cd apps/order/rpc && go run order.go

product:
	cd apps/product/rpc && go run product.go

reply:
	cd apps/reply/rpc && go run reply.go

test_order:
	curl http://127.0.0.1:8888/v1/order/list?uid=123

test_productdetail:
	curl http://127.0.0.1:8888/v1/product/detail?product_id=1

rpc[reply]:
	cd apps/reply/rpc && goctl rpc protoc reply.proto --go_out=. --go-grpc_out=. --zrpc_out=.

rpc[product]:
	cd apps/product/rpc && goctl rpc protoc product.proto --go_out=. --go-grpc_out=. --zrpc_out=.

rpc[order]:
	cd apps/order/rpc && goctl rpc protoc order.proto --go_out=. --go-grpc_out=. --zrpc_out=.

api[app]:
	cd apps/app/api && goctl api go -api api.api -dir .